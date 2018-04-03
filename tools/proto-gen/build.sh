#!/bin/sh
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

# build proto msg
SERVER_PATH=${DIR}/../../server/src/golf/comm/msg
CLIENT_PATH=${DIR}/../../client/common/msg
protoc --go_out=$SERVER_PATH --js_out=$CLIENT_PATH protos/*.proto

# build proto mgr
${DIR}/tool-proto-gen/build.sh 