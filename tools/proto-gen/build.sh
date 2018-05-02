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
INTERNAL_OUT=./build/internal

mkdir -p $SERVER_OUT
mkdir -p $CLIENT_OUT
mkdir -p $FAIRY_OUT
mkdir -p $INTERNAL_OUT

PLUGIN_PATH=./tool-proto-gen/protoc-gen-fairy
./bin/protoc --plugin=./bin/protoc-gen-go --go_out=$SERVER_OUT --js_out=import_style=commonjs,binary:$CLIENT_OUT --plugin=$PLUGIN_PATH --fairy_out="js,go;msgid=1000":$FAIRY_OUT  --proto_path=protos protos/*.proto
./bin/protoc --plugin=./bin/protoc-gen-go --go_out=$SERVER_OUT --plugin=$PLUGIN_PATH --fairy_out="go;msgid_open=false":$INTERNAL_OUT -I protos -I protos_internal protos_internal/*.proto

#copy server
SERVER_MSG_ROOT=../../server/src/server/comm/msg
mkdir -p $SERVER_MSG_ROOT
# rm -f $SERVER_MSG_ROOT

cp -f ./build/server/*.go   $SERVER_MSG_ROOT
cp -f ./build/internal/msg.go $SERVER_MSG_ROOT/msg_internal.go
cp -f ./build/fairy/msg.go  $SERVER_MSG_ROOT

#copy client