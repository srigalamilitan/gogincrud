FROM golang:1.14 as builder
LABEL maintener="krisna putra <anakgembala21@gmail.com>"
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gogincrud .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/ .

CMD ["./gogincrud"]