FROM golang:1.17-alpine

# 下記を参考にマルチステージングビルドに書き換え
# https://qiita.com/gs1068/items/ed1637cf57400ed6ca21

WORKDIR /go/src/app
COPY ./app .

RUN apk update && \
    apk add curl && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh
    # go get github.com/go-sql-driver/mysql@latest

#RUN go get github.com/jinzhu/gorm

# RUN go get -tags 'mysql' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
# RUN go get github.com/golang-migrate/migrate/v4/cmd/migrate

#RUN go get github.com/spf13/viper
#RUN go get -u github.com/o-sql-driver/mysql

# oxequa/realizeが不要かも
#RUN go get -u github.com/oxequa/realize \
    # sqlを使うためのモジュール
#    && go get github.com/go-sql-driver/mysql

EXPOSE 8080

CMD ["fresh"]
