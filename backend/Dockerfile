FROM golang:1.22.5
RUN  mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o main 
CMD ["/app/main"]