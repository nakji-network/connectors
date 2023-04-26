FROM rust:1.68.2

WORKDIR /sui

RUN apt-get update && apt-get -y install cmake clang

# Build dependencies
COPY Cargo.toml .
COPY Cargo.lock .
RUN mkdir src && echo "// dummy file" > src/lib.rs
RUN cargo build --release
RUN rm -r src

# Build application
COPY . .
RUN cargo build --release
RUN mv target/release/sui .

CMD ["sui"]