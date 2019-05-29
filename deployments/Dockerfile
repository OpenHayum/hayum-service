ARG GOLANG_VERSION="latest"

FROM golang:${GOLANG_VERSION}

ARG ENV="development"
ENV GO_ENV=${ENV}
ENV GO111MODULE=on

LABEL maintainer="Devajit Asem"

WORKDIR $GOPATH/src/hayum/core_apis

COPY . .

RUN go install ./cmd/hayum

EXPOSE 8080

CMD ["hayum"]