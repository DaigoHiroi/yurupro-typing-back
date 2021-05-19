FROM golang:1.15.2-alpine

WORKDIR /go/src/app
COPY ./app .

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh && \
    go get github.com/go-sql-driver/mysql

# oxequa/realizeが不要かも
#RUN go get -u github.com/oxequa/realize \
    # sqlを使うためのモジュール
#    && go get github.com/go-sql-driver/mysql

EXPOSE 8080

CMD ["fresh"]
