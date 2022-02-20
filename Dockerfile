FROM golang:1.17

WORKDIR /goApp
COPY . .

RUN apt update
RUN apt install sqlite3
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 9099

CMD ["willette_api"]
