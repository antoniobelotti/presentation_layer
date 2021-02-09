FROM golang:1.15-alpine3.12 AS builder-container

ENV APP_USER app
ENV APP_HOME /go/src/presentation_layer
ENV APP_GROUP app_user_group
ENV APP_USER app_user

RUN addgroup -S $APP_GROUP && adduser -S $APP_USER -G $APP_GROUP
RUN mkdir -p $APP_HOME && chown -R $APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER

COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o app

FROM alpine:3.13.1

ENV APP_HOME /go/src/presentation_layer
ENV APP_GROUP app_user_group
ENV APP_USER app

RUN addgroup -S $APP_GROUP && adduser -S $APP_USER -G $APP_GROUP
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY templates/ templates/
COPY --chown=0:0 --from=builder-container $APP_HOME/app $APP_HOME

EXPOSE 80
USER $APP_USER
CMD ["./app"]