FROM public.ecr.aws/docker/library/golang:1.17-bullseye as build

WORKDIR /build
COPY . /build

RUN CGO_ENABLED=0 go build -o balance-server cmd/*.go

FROM public.ecr.aws/docker/library/alpine:3

RUN adduser -S -D -H appuser
USER appuser

COPY --from=build /build/balance-server /app/

WORKDIR /app

ENTRYPOINT ["./balance-server"]
