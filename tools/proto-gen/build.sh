#!/bin/sh
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

cd $DIR

REBUILD=$1

# rebuild msgid
if [ $REBUILD ]; then
    echo "rebuild msgid"
    rm -f ./msgid.json
fi 

# build proto msg
# SERVER_PATH=../../server/src/golf/comm/msg
# CLIENT_PATH=../../client/common/msg
SERVER_OUT=./build/server
CLIENT_OUT=./build/client
FAIRY_OUT=./build/fairy

mkdir -p $SERVER_OUT
mkdir -p $CLIENT_OUT
mkdir -p $FAIRY_OUT

PLUGIN_PATH=./tool-proto-gen/protoc-gen-fairy
protoc  --go_out=$SERVER_OUT --js_out=import_style=commonjs,binary:$CLIENT_OUT --plugin=$PLUGIN_PATH --fairy_out=js,go:$FAIRY_OUT  --proto_path=protos protos/*.proto

# copy fairy