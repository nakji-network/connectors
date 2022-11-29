FROM blepai/golang-proto:latest AS base

RUN apk add --no-cache build-base linux-headers

# download dependencies first
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

FROM base as build
# build. CGO_ENABLED=1 required for conflient-kafka-go
WORKDIR /src
RUN mkdir /src/bin
RUN CGO_ENABLED=1 GOFLAGS="-tags=musl" go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o bin ./...

FROM build as test
ENV CGO_ENABLED=1 
ENV GOFLAGS="-tags=musl"
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
CMD ["go", "test", "./..."]

#FROM scratch AS final
# can't use scratch for confluent-kafka-go
FROM alpine:3.15 AS final
RUN apk add --no-cache

ARG PACKAGE=./usr/local
RUN echo ${PACKAGE}

COPY --from=build /src/bin /usr/local/bin
COPY --from=build /src/manifest.yaml /
CMD ["ash"]
