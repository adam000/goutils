#!/usr/bin/env bash

set +e

VERSION="release-branch.go1.13"

export CGO_ENABLED=0

ROOT=$HOME/go
mkdir -p $ROOT
cd $ROOT

# Versions later than go 1.4 need go 1.4 to build successfully.
if [[ ! -d go1.4 ]]; then
    git clone --branch release-branch.go1.4 --depth=1 git@github.com:golang/go.git go1.4
    cd go1.4/src
    ./make.bash
    cd ../..
fi

# Make sure we build the new version from scratch
if [[ -d "$VERSION" ]]; then
    rm -rf $ROOT/$VERSION
fi

git clone --branch $VERSION --depth=1 git@github.com:golang/go.git $VERSION

cd $VERSION/src
GOROOT_BOOTSTRAP=$ROOT/go1.4 ./make.bash
cd ../..


[[ -d src ]] && rm src bin pkg
ln -s $VERSION/src src
ln -s $VERSION/bin bin
ln -s $VERSION/pkg pkg

