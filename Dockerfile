FROM golang:1.17

WORKDIR /go/src/andrewwillette
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["willette_api"]