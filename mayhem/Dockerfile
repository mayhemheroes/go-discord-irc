FROM golang:1.18 as builder

COPY . /go-discord-irc
WORKDIR /go-discord-irc/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang


RUN go-fuzz-build -libfuzzer -o fuzzParse.a && \
    clang -fsanitize=fuzzer fuzzParse.a -o fuzzParse.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /go-discord-irc/mayhem/fuzzParse.libfuzzer /