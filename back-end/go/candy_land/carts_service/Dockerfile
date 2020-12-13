FROM alpine:3.12.0
RUN apk --no-cache add ca-certificates

WORKDIR /root/

ADD dist/main .

CMD ["./main"]