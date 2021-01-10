#
# @file    CMakeLists.txt
# @copyright defined in meycoin/LICENSE.txt
#

cmake_minimum_required(VERSION 3.0)

project(meycoin NONE)

set(BIN_DIR ${CMAKE_CURRENT_LIST_DIR}/bin)
set(PROTO_DIR ${CMAKE_CURRENT_LIST_DIR}/mey-protobuf/proto)

if(CMAKE_BUILD_TYPE MATCHES "Debug")
    set(GCFLAGS "-gcflags=\"-N" "-l\"")
    set(GFLAG -tags Debug)
endif()

add_custom_target(build ALL DEPENDS meycoincli meycoinsvr meycoinluac brick)

add_custom_target(meycoincli GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} -ldflags \"-X github.com/meeypioneer/meycoin/cmd/meycoincli/cmd.githash=`git describe --tags`\" ./cmd/meycoincli/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS libtool)

add_custom_target(meycoinsvr GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} -ldflags \"-X main.githash=`git describe --tags` -X main.gitRevision=`git rev-parse --short HEAD` -X main.gitBranch=`git rev-parse --symbolic-full-name --abbrev-ref HEAD`\" ./cmd/meycoinsvr/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS libtool)

add_custom_target(wezen GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} -ldflags \"-X main.githash=`git describe --tags`\" ./cmd/wezen/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})

add_custom_target(wetool GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} -ldflags \"-X github.com/meeypioneer/meycoin/cmd/wetool/cmd.githash=`git describe --tags`\" ./cmd/wetool/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})

add_custom_target(meycoinluac GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} -ldflags \"-X main.githash=`git describe --tags`\" ./cmd/meycoinluac/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS libtool)

add_custom_target(brick GO111MODULE=on GOBIN=${BIN_DIR} go install ${GCFLAGS} ${GFLAG} -ldflags \"-X 'github.com/meeypioneer/meycoin/cmd/brick/context.GitHash=`git describe --tags`'
-X 'github.com/meeypioneer/mey-library/log.defaultConfStr=`cat ./cmd/brick/arglog.toml`'\"  ./cmd/brick/...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS libtool)

add_custom_target(deps DEPENDS libtool)

add_custom_target(check GO111MODULE=on go test ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS build)
add_custom_target(cover-check GO111MODULE=on go test -coverprofile c.out ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
    DEPENDS build)

add_custom_target(distclean GO111MODULE=on go clean ..
	COMMAND rm -rf ${BIN_DIR}/meycoin* ${BIN_DIR}/brick
	DEPENDS libtool-clean)

add_custom_target(protoc
	COMMAND protoc -I/usr/local/include -I${PROTO_DIR} --go_out=plugins=grpc:${CMAKE_CURRENT_LIST_DIR}/types ${PROTO_DIR}/*.proto
	COMMAND GO111MODULE=on go build ../types/...)

add_custom_target(protoclean rm -f ../types/*.pb.go)


add_subdirectory(libtool)
