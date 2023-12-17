FROM golang:1.20

WORKDIR /usr/src/Backend

#For caching the layers only go.mod and go.sum are initially copied
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o MiniRoute .

EXPOSE 8080

CMD [ "./MiniRoute" ]