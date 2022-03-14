FROM golang:1.17 AS builder

ARG APP_NAME
ARG APP_TYPE
ARG APP_ENV

COPY . /src
WORKDIR /src/app/${APP_NAME}

# RUN export GOPROXY=https://goproxy.cn  make clean && make build
RUN go env -w GOPROXY=https://goproxy.cn  && make clean && make build

# ==============================================================

FROM debian:stable-slim

ARG APP_NAME
ARG APP_TYPE
ARG APP_ENV
ENV MO_APP_NAME urlshorten.${APP_TYPE}.${APP_NAME}
ENV MO_APP_ENV ${APP_ENV}

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_NAME}/bin /app
COPY --from=builder /src/app/${APP_NAME}/config/${APP_ENV}/config.yaml /app/config/config.yaml

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server", "-conf", "./config"]
