package base

import (
	"app/app/base/messages"
	ci18n "app/config/i18n"
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
)

var (
	Galidator = galidator.New().CustomMessages(galidator.Messages{
		"required": "$field is required",
		"min":      "$field must be higher equal to rule",
	})
)

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`([a-z_0-9])([A-Z])`)

type conventionalMarshallerFromPascal struct {
	Value any
}

func convertToCamelCase(marshalled []byte) []byte {
	return keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			if len(match) > 2 {
				converted := bytes.ToLower(wordBarrierRegex.ReplaceAll(
					match,
					[]byte(`${1}_${2}`),
				))

				var result []byte
				underscore := false
				for i := 1; i < len(converted)-1; i++ {
					if converted[i] == '_' {
						underscore = true
					} else {
						if underscore {
							result = append(result, byte(unicode.ToUpper(rune(converted[i]))))
							underscore = false
						} else {
							result = append(result, converted[i])
						}
					}
				}
				result = append([]byte{converted[0]}, result...)
				result = append(result, converted[len(converted)-1])
				return result
			}
			return match
		},
	)
}

func (c conventionalMarshallerFromPascal) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	if err != nil {
		return nil, err
	}
	naming := viper.GetString("HTTP_JSON_NAMING")

	val := reflect.TypeOf(c.Value)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {

		field, ok := val.FieldByName("json")
		if ok {

			if field.Tag.Get("naming") != "" {
				naming = field.Tag.Get("naming")
			}
		}
	}

	var converted []byte
	switch naming {
	case "snake_case":

		converted = keyMatchRegex.ReplaceAllFunc(
			marshalled,
			func(match []byte) []byte {
				return bytes.ToLower(wordBarrierRegex.ReplaceAll(
					match,
					[]byte(`${1}_${2}`),
				))
			},
		)
	case "camel_case":
		converted = convertToCamelCase(marshalled)
	case "pascal_case":
		return marshalled, nil
	default:
		return nil, err
	}

	return converted, nil
}

func defaultJSON(ctx *gin.Context, code int, msgID string, data any, paginate *ResponsePaginate, params ...map[string]string) error {
	localizer := i18n.NewLocalizer(ci18n.Bundle, ctx.GetHeader("Accept-Language"))

	var param any
	if len(params) > 0 {
		param = params[0]
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: msgID, TemplateData: param})

	if err != nil || msg == "" {
		ctx.JSON(code, Response[any]{
			ResponseStatus: &ResponseStatus{
				Message: msgID,
				Code:    code,
			},
			Data:     data,
			Paginate: paginate,
		})
		return nil
	}

	ctx.JSON(code, conventionalMarshallerFromPascal{Response[any]{
		ResponseStatus: &ResponseStatus{
			Message: msg,
			Code:    code,
		},
		Data:     data,
		Paginate: paginate,
	}})
	return nil
}

func Success(ctx *gin.Context, data any) error {
	return defaultJSON(ctx, http.StatusOK, messages.Success, data, nil)
}

func Paginate(ctx *gin.Context, data any, page *ResponsePaginate) error {
	return defaultJSON(ctx, http.StatusOK, messages.Success, data, page)
}

func BadRequest(ctx *gin.Context, message string, data any, params ...map[string]string) error {
	return defaultJSON(ctx, http.StatusBadRequest, message, data, nil)
}

func Unauthorized(ctx *gin.Context, message string, data any, params ...map[string]string) error {
	return defaultJSON(ctx, http.StatusUnauthorized, message, data, nil)
}

func Forbidden(ctx *gin.Context, message string, data any, params ...map[string]string) error {
	return defaultJSON(ctx, http.StatusForbidden, message, data, nil)
}

func InternalServerError(ctx *gin.Context, message string, data any, params ...map[string]string) error {
	return defaultJSON(ctx, http.StatusInternalServerError, message, data, nil)
}
