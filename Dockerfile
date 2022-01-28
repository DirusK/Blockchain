# first (build) stage
FROM golang:alpine as builder
ARG PROJECT_NAME="blockchain"

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 && go build -o cmd/$PROJECT_NAME cmd/main.go


# final (target) stage
FROM alpine:latest

LABEL author="insomnia" \
      version="v1.0.0"

EXPOSE 5000

WORKDIR /app

COPY --from=builder /app/cmd/$PROJECT_NAME ./
COPY --from=builder /app/cmd/config.tmpl.yaml ./config.yaml

ENTRYPOINT [ "./blockchain", "-c", "config.yaml" ]
