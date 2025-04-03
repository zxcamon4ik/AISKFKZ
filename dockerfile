FROM golang:1.24.1-alpine

ARG USER=appuser
ARG WORKDIR=/wd

# Add bash, git and required tools (apk not apt!)
RUN apk add --no-cache bash git

# Create user + set workdir
RUN adduser -D -g '' ${USER} && \
    mkdir -p ${WORKDIR} && \
    chown -R ${USER}:${USER} ${WORKDIR}

WORKDIR ${WORKDIR}

COPY --chown=${USER}:${USER} go.mod go.sum ./
RUN go mod download

COPY --chown=${USER}:${USER} ./api ./api
COPY --chown=${USER}:${USER} ./cmd ./cmd
COPY --chown=${USER}:${USER} ./configs ./configs
COPY --chown=${USER}:${USER} ./docker/service/entrypoint.sh ./entrypoint.sh

RUN chmod +x entrypoint.sh

USER ${USER}

RUN go build -o main ./cmd/main.go

ENTRYPOINT ["./entrypoint.sh"]
CMD ["./main"]
