#!/bin/bash
set -e 

echo "#### backend building start"

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
APP_DIR=$SCRIPTPATH/../dist/

cd $SCRIPTPATH

if ! [ -d $APP_DIR ]; then
  mkdir -v $APP_DIR
fi

go build -v -o $SCRIPTPATH/app.exe .

mv app.exe $APP_DIR

cp -r templates $APP_DIR

echo "#### backend building finish"
