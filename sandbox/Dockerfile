FROM golang:1.12.6 as tests

WORKDIR /src

ADD . .

RUN go test ./...


FROM golang:1.12.6 as builder

WORKDIR /src

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/go-app


FROM scratch

WORKDIR /root

COPY --from=buildr /dist .

CMD [ "./go-app" ]
