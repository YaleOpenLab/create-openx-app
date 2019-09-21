#!/bin/sh
mkdir $1
cp -r .template/* $1
cd $1
echo "Replacing YaleOpenLab/opensolar with $1/$2"
find . -name '*.go' -print0 | xargs -0 sed -i "" "s,github.com/org/plat,github.com/$1/$2,g"
cd ..
echo $GOPATH
