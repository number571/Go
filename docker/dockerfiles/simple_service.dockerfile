FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget
RUN wget https://dl.google.com/go/go1.20.3.linux-amd64.tar.gz && \ 
    tar -C /opt -xzf go1.20.3.linux-amd64.tar.gz

WORKDIR /app
ENV PATH="${PATH}:/opt/go/bin"
COPY ./go.* ./app/* /app/
RUN go build -o service . 

ENTRYPOINT ["./service"]
CMD [":80"]
