#!/bin/sh

mkdir $1
cp -r .template/* $1
cd $1
# echo "Replacing YaleOpenLab/opensolar with $1/$2"
find . -name '*.go' -print0 | xargs -0 sed -i "" "s,github.com/YaleOpenLab/create-openx-app/.template,github.com/$1/$2,g"
yes="y"

# echo $3, $4, $6

if [ "$3" == "$yes" ] ; then
  echo "Adding extra options for investors"
  cd core
  rm investor.go
  mv investor_voter.gotemp investor.go
else
  echo "Deleting extra options for investors"
  find . -name '*investor_voter.gotemp' -exec rm "{}" \;
fi

if [ "$4" == "$yes" ] ; then
  # ie the person wants to not omit the handlers
  echo "Adding extra options for recipients"
else
  echo "Deleting extra options for recipients"
  find . -name '*recipient_options.go' -exec rm "{}" \;
fi

cd notif
find . -name '*.go' -print0 | xargs -0 sed -i "" "s,REPLACEME,$5,g"
cd ..

if [ "$6" == "$yes" ] ; then
  echo "Adding support for other blockchains"
else
  echo "Omitting other blockchain handlers"
  cd core
  rm project.go contract.go
  mv project_stellar.gotemp project.go
  mv contract_stellar.gotemp contract.go
fi

cd ..
# cd ..
# echo $GOPATH
# mv $1 $GOPATH/
# cd $GOPATH/$1
# go build
# echo "Template moved under GOPATH, please build form there to see if the platform works as expected"
