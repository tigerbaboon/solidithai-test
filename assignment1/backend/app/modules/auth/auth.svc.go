package auth

import (
	"app/app/base/messages"
	"app/app/modules/accesstoken"
	authdto "app/app/modules/auth/dto"
	"app/app/modules/user"
	"app/app/util/hashing"
	"app/app/util/jwt"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type AuthService struct {
	db          *bun.DB
	user        *user.UserModule
	accesstoken *accesstoken.AccessTokenModule
}

func newService(db *bun.DB, user *user.UserModule, accesstoken *accesstoken.AccessTokenModule) *AuthService {
	return &AuthService{
		db:          db,
		user:        user,
		accesstoken: accesstoken,
	}
}

func (s *AuthService) Login(ctx context.Context, req authdto.LoginRequest) (*string, bool, error) {

	user, err := s.user.Svc.ExistByEmail(ctx, req.Email)
	if err != nil {
		return nil, false, err
	}

	if user == nil {
		return nil, true, errors.New(messages.UserOrPasswordIncorrect)
	}

	userPass, err := s.user.Svc.ExistPassword(ctx, user.ID)
	if err != nil {
		return nil, false, err
	}

	if userPass == nil {
		return nil, true, errors.New(messages.UserPasswordNotFound)
	}

	if !hashing.CheckPasswordHash([]byte(userPass.Password), []byte(req.Password)) {
		return nil, true, errors.New(messages.UserOrPasswordIncorrect)
	}

	claim := jwt.ClaimData{
		UserID: user.ID,
	}

	accesstoken, err := jwt.CreateToken(claim)
	if err != nil {
		return nil, false, err
	}

	err = s.accesstoken.Svc.Create(ctx, user.ID, *accesstoken)

	return accesstoken, false, err
}
