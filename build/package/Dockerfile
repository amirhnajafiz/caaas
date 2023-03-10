FROM golang:1.19-alpine

LABEL app="authx"
LABEL docker_file_version="v0.1.0"
LABEL database_driver="mysql"

MAINTAINER amirhossein.najafizadeh21@gmail.com

ENV VERSION="v0.1.0"
ENV APP="authx"

COPY . .

RUN go build -o main
RUN chmod +x ./main
RUN ./main migrate

EXPOSE 8080

CMD ./main http
