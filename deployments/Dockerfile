ARG GOLANG_VERSION="latest"

FROM golang:${GOLANG_VERSION}

ARG ENV="development"
ENV GO_ENV=${ENV}

LABEL maintainer="Devajit Asem"

WORKDIR $GOPATH/src/hayum/core_apis

COPY . .

# Get dependencies
RUN GO111MODULE=on go get -d ./cmd/hayum...
# Build hayum
RUN go install ./cmd/hayum

EXPOSE 8080

CMD ["hayum"]