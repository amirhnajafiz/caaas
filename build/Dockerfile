# using golang:1.22 base image
FROM golang:1.22-alpine

# don't change workdir
WORKDIR /usr/share/caass

COPY internal/ internal/
COPY pkg/ pkg/
COPY main.go main.go

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o main && chmod +x ./main

CMD [ "./main" ]
