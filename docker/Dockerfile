FROM frolvlad/alpine-glibc:latest

ARG APP
WORKDIR /app
ADD ${APP} /app
ENTRYPOINT ./main
