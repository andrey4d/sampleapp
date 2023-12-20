FROM golang:1.21-alpine as builder

COPY . /src
WORKDIR /src
RUN apk add make
RUN make build
RUN cp -R web bin/


FROM scratch

LABEL app="sampleapp"
LABEL maintainer="andrey4d.dev@gmial.com"

COPY --from=builder /src/bin /opt
WORKDIR /opt
CMD [ "/opt/backend" ]