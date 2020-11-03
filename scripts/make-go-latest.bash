#!/usr/bin/env bash

set +e

if ! which -s gcc; then
    echo "gcc needs to be installed, but I couldn't find it. Bailing..."
    exit 0
fi

VERSION="release-branch.go1.15"

export CGO_ENABLED=0

ROOT=$HOME/go
mkdir -p $ROOT
cd $ROOT

# Versions later than go 1.4 need go 1.4 to build successfully.
if [[ ! -d go1.4 ]]; then
    git clone --branch release-branch.go1.4 --depth=1 https://github.com/golang/go.git go1.4
    cd go1.4/src
    ./make.bash
    cd ../..
fi

# Make sure we build the new version from scratch
if [[ -d "$VERSION" ]]; then
    rm -rf $ROOT/$VERSION
fi

git clone --branch $VERSION --depth=1 https://github.com/golang/go.git $VERSION

cd $VERSION/src
GOROOT_BOOTSTRAP=$ROOT/go1.4 ./make.bash
cd ../..


# Link the latest versions into the parent directory for easy access
[[ -d src ]] && rm src bin pkg
ln -s $VERSION/src src
ln -s $VERSION/bin bin
ln -s $VERSION/pkg pkg
ln -s $VERSION/misc misc

# Make sure that ~/go/bin is in $PATH at least for this shell so that vim -c works
if [[ ":$PATH:" != *":$HOME/go/bin:"* ]]; then
    export PATH=$PATH:$HOME/go/bin
fi

# Install or update the latest Go plugins
vim -c "UpdateAddons"
vim -c "GoInstallBinaries"
vim -c "GoUpdateBinaries"
