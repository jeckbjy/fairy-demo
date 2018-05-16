#!/bin/sh
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

${DIR}/tool-table-gen/bin/build.sh -c ./conf.yaml

