FROM node:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN rm -rf node_modules/

RUN npm i

ADD . .

RUN npm run build

ENV PORT="5000"

FROM node:alpine

WORKDIR /root/

RUN mkdir src/

COPY --from=builder app/dist ./src
COPY --from=builder app/.env .
COPY --from=builder app/package.json .
COPY --from=builder app/src/schema.graphql src/

RUN npm i

EXPOSE 4000

CMD node src/index.js 