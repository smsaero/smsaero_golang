FROM golang:latest AS build

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/smsaero_send demo/main.go


FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/smsaero_send /usr/bin/smsaero_send

CMD ["/usr/bin/smsaero_send"]
