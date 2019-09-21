#!/bin/sh

mkdir $1
cp -r .template/* $1
cd $1
echo "Replacing YaleOpenLab/opensolar with $1/$2"
find . -name '*.go' -print0 | xargs -0 sed -i "" "s,github.com/YaleOpenLab/create-openx-app/.template,github.com/$1/$2,g"
investor="investor"
recipient="recipient"

echo $3, $4

if [ "$3" == "$investor" ] ; then
  echo "Adding extra options for investors"
else
  echo "Deleting extra options for investors"
  find . -name '*investor_options.go' -exec rm "{}" \;
fi

if [ "$4" == "$recipient" ] ; then
  echo "Adding extra options for recipients"
else
  echo "Deleting extra options for recipients"
  find . -name '*recipient_options.go' -exec rm "{}" \;
fi

# cd ..
# echo $GOPATH
# mv $1 $GOPATH/
# cd $GOPATH/$1
# go build
# echo "Template moved under GOPATH, please build form there to see if the platform works as expected"
