ARG GOLANG_VERSION="latest"
FROM golang:${GOLANG_VERSION}

ARG ENV="development"
ENV GO_ENV=${ENV}
LABEL maintainer="Devajit Asem"
WORKDIR $GOPATH/src/hayum/core_apis
COPY . .
RUN go get -d ./...
RUN go install ./...
EXPOSE 8080

CMD ["core_apis"]