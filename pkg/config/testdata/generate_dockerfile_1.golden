FROM golang:1.15.2
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /src
COPY . .

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o /app" -command="/app"