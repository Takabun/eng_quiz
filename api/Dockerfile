FROM golang:1.13
# FROM golang:onbuild
# FROM golang:1.3.1-onbuild

# RUN go get github.com/labstack/echo/...

# Docker上の作業ディレクトリ設定
WORKDIR ./go 
# ローカルのファイルをdocker にコピー
# ADD . 
COPY . /go

CMD ["go", "run", "main.go"]