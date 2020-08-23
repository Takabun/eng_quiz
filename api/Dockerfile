FROM golang:1.13

RUN go get github.com/labstack/echo && \
    go get github.com/jinzhu/gorm  && \
    go get github.com/dgrijalva/jwt-go  && \
    go get github.com/go-sql-driver/mysql

# WORKDIR ./go  Docker上の作業ディレクトリ設定(不要。pwdが/go/goになる)
# ADD . # ローカルのファイルをdocker にコピー

#元はこうだったCOPY . /go
COPY . .

CMD ["go", "run", "main.go"]