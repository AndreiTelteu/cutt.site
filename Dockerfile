FROM golang:1.23
ARG USER_ID
ARG GROUP_ID

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

RUN groupadd -r app --gid $GROUP_ID && \
    useradd -u $USER_ID -r -g app -m -d /home/app -s /usr/bin/bash -c "App user" app && \
    chown -R app:app /app && \
    mkdir /go-cache && chown -R app:app /go-cache /go

USER app
RUN go env -w GOCACHE=/go-cache/go-cache && go env -w GOMODCACHE=/go-cache/gomod-cache

EXPOSE 3000
CMD ["air", "-c", ".air.toml"]

# for debuging
# CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"
