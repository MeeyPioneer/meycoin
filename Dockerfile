FROM golang:1.12.5-alpine3.9 as builder
RUN apk update && apk add git cmake build-base m4
COPY . meycoin
RUN cd meycoin && make meycoinsvr

FROM alpine:3.9
RUN apk add libgcc
COPY --from=builder /go/meycoin/bin/meycoinsvr /usr/local/bin/
COPY --from=builder /go/meycoin/libtool/lib/* /usr/local/lib/
COPY --from=builder /go/meycoin/Docker/conf/* /meycoin/
ENV LD_LIBRARY_PATH="/usr/local/lib:${LD_LIBRARY_PATH}"
WORKDIR /meycoin/
CMD ["meycoinsvr", "--home", "/meycoin"]
EXPOSE 7845 7846 6060
