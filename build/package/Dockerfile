# from golang alpine image
FROM golang:1.19-alpine

# labels for swarm and other managers
LABEL app="authx"
LABEL docker_file_version="v0.1.0"
LABEL database_driver="mysql"

# maintainer
MAINTAINER amirhossein.najafizadeh21@gmail.com

# env variables
ENV VERSION="v0.1.0"
ENV APP="authx"

# copy all files
COPY . .

# building go executable file
RUN go build -o main
RUN chmod +x ./main
# migrate database
RUN ./main migrate

# expose port 8080
EXPOSE 8080

# start http service
CMD ./main http
