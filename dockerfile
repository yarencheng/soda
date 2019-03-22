FROM golang:1.11 as builder

COPY . /src
WORKDIR /src/go_src
RUN go install

FROM ubuntu:18.04

COPY --from=builder /go/bin/go_src /usr/bin/soda
ENV PHOTO_DIR=/photos
ENTRYPOINT /usr/bin/soda