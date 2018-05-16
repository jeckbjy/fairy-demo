#!/bin/sh
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

pushd $DIR > /dev/null
./tool-table-gen/bin/build.sh
popd > /dev/null
