FROM clux/muslrust:stable

RUN apt update && apt upgrade -y
RUN apt install -y protobuf-compiler libprotobuf-dev

WORKDIR /usr/src/app

CMD ["cargo", "build", "--release"]