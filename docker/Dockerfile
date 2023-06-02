FROM ubuntu:20.04

RUN apt-get update -y
RUN apt-get install -y wget
RUN rm /bin/sh && ln -s /bin/bash /bin/sh 
RUN wget https://dl.google.com/go/go1.20.3.linux-amd64.tar.gz
RUN tar -C /opt -xzf go1.20.3.linux-amd64.tar.gz

WORKDIR /app
ENV PATH="${PATH}:/opt/go/bin"
COPY ./go.* /app
COPY ./app /app
RUN go build -o service . 

ENTRYPOINT ["./service"]
CMD [":80"]
