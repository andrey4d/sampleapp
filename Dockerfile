FROM golang:1.21-alpine3.18 as builder

COPY . /src
WORKDIR /src
RUN apk add make
RUN make build


FROM scratch

LABEL app="sampleapp"
LABEL maintainer="andrey4d.dev@gmial.com"

COPY --from=builder /src/bin /

CMD ["/backend"]