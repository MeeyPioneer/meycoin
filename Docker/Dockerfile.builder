FROM golang:1.12.5-alpine3.9
ARG GIT_TAG=master
RUN apk update && apk add git cmake build-base m4
RUN git clone --branch ${GIT_TAG} --recursive https://github.com/meeypioneer/meycoin.git \
    && cd meycoin \
    && make meycoinsvr wezen wetool meycoincli meycoinluac brick

