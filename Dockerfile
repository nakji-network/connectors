FROM blepai/golang-proto:latest AS base

RUN apk add --no-cache build-base linux-headers

# download dependencies first
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# build protobufs
RUN protoc --proto_path=. \
                --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional $(find . -iname "*.proto")

FROM base as build
# build. CGO_ENABLED=1 required for conflient-kafka-go
WORKDIR /src
RUN mkdir /src/bin
RUN CGO_ENABLED=1 GOFLAGS="-tags=musl" go build -o bin ./...

FROM build as test
ENV CGO_ENABLED=1 
ENV GOFLAGS="-tags=musl" 
CMD ["go", "test", "./..."]

#FROM scratch AS final
# can't use scratch for confluent-kafka-go
FROM alpine:3.15 AS final
RUN apk add --no-cache protoc protobuf-dev

ARG PACKAGE=./usr/local
RUN echo ${PACKAGE}

COPY --from=build /src/bin /usr/local/bin
COPY --from=build /src/manifest.yaml /
# Copy .proto files
COPY --from=build /src ${PACKAGE}
RUN find ${PACKAGE} -type f -not -iname "*.proto" -delete \
 && find ${PACKAGE} -type f -iname "*.proto" -exec sh -c 'mv "$1" "${1%.proto}.proto.tmp"' _ {} \; \
 && mkdir ./proto-tmp \
 && mv ${PACKAGE} ./proto-tmp \
 && chmod -R 777 proto-tmp

CMD ["ash"]
