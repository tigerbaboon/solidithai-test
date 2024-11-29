FROM node:20-alpine as build

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./

RUN yarn

COPY . ./

RUN yarn build

FROM node:20-alpine as serve

WORKDIR /app
COPY --from=build /app/.output ./

EXPOSE 8080

ENTRYPOINT ["node","/app/server/index.mjs"]
