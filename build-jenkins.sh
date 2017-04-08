#!/bin/sh

export PATH=$PATH:$GOPATH/bin
export PRJ=`git config --get remote.origin.url | sed 's/^https:\/\///' | sed 's/\.git$//'`
export BASE_DIR=`pwd`
export BUILD_DIR=build

# Run fresh
rm -rf $BUILD_DIR

# Install gometalinter
if [ ! -x "$GOPATH/bin/gometalinter" ]; then
  go get -u github.com/alecthomas/gometalinter
  gometalinter --install
fi

# Copy source file
mkdir -p $BUILD_DIR/src/$PRJ
ls -1 | grep -v ^$BUILD_DIR | xargs -I{} cp -pr {} $BUILD_DIR/src/$PRJ/

cd $BUILD_DIR/src/$PRJ
GOPATH=$BASE_DIR/$BUILD_DIR

go get -v ./...
gometalinter ./...
go test -v ./...
go install -v ./...
