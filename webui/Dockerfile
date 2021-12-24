FROM golang:1.17
RUN apt-get update
RUN apt-get install -y git python jq curl

RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN npm install gulp -g
RUN npm install yarn -g

WORKDIR $GOPATH/src/github.com/moshebe/gebug
ENV VUE_APP_PORT 3030
COPY go.mod go.sum ./
COPY ./pkg ./pkg
RUN go mod download

COPY ./webui ./webui
WORKDIR $GOPATH/src/github.com/moshebe/gebug/webui/frontend
RUN yarn install
RUN yarn build && cp -r src/assets dist

WORKDIR $GOPATH/src/github.com/moshebe/gebug/webui
RUN cd frontend && cp -r src/assets dist
ENTRYPOINT go run server.go