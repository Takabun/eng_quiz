FROM golang:1.14

RUN go get github.com/labstack/echo && \
    go get github.com/labstack/echo/middleware && \
    go get github.com/jinzhu/gorm  && \
    go get github.com/dgrijalva/jwt-go  && \
    go get github.com/go-sql-driver/mysql

COPY . .

CMD ["go", "run", "main.go"]