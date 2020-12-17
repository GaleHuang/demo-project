#!/bin/bash

echo "going to build proto..."
protoc -I=./ --go_out=plugins=grpc:../ ./*.proto
echo "build complete...."


