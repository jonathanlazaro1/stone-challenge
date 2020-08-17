FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
RUN chmod +x ./cmd/docker/docker-start.sh
ENTRYPOINT [ "./cmd/docker/docker-start.sh" ]