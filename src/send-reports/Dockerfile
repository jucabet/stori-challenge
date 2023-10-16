FROM public.ecr.aws/docker/library/golang:1.19-alpine3.15 as builder

ARG GO_MODULES_TOKEN=token

RUN apk update

RUN apk add git
RUN apk add build-base

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=arm64 go build -tags musl -tags dynamic cmd/main.go 

# Run the Go Binary in Alpine.
FROM public.ecr.aws/docker/library/alpine:3.18

RUN apk update
RUN apk add build-base

RUN addgroup -S appgroup && adduser -S appuser -G appgroup --uid 10001 

WORKDIR /app
COPY --from=builder  app/main main
RUN chmod +x ./main
USER appuser
EXPOSE 8080

CMD ["./main"]