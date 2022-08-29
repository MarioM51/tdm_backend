#!/bin/bash
set -e 

echo "=== backend building start"

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

env=${env:-null}
dist_directory=${dist_directory:-null}
exe_filename=${dist_directory:-null}

while [ $# -gt 0 ]; do
   if [[ $1 == *"--"* ]]; then
        param="${1/--/}"
        declare $param="$2"
        # echo $1 $2 // Optional to see the parameter:value result
   fi
  shift
done

if [ "$env" == "null" ]; then
  echo "--env: enviroment required"
  exit 1
fi

if [ "$exe_filename" == "null" ]; then
  echo "--exe_filename: name of .exe result"
  exit 1
fi

if [ "$dist_directory" == "null" ]; then
  echo "--dist_directory: distibution directory path required"
  exit 1
fi


if ! [ -d $APP_DIR ]; then
  mkdir -v $APP_DIR
fi


if [ $env == "test" ]; then
  echo "====== for linux"

  docker run --rm \
    --name go_builder \
    -v /$SCRIPTPATH:/src \
    -v /$dist_directory:/dist \
    -v /$SCRIPTPATH/pkg_lin:/go/pkg \
    -w //src \
    golang:1.18.5 \
    env CGO_ENABLED=0 go build -v -o //dist/$exe_filename .

fi

if [ $env == "dev" ]; then
  echo "====== for windows"
  go build -v -o $dist_directory/$exe_filename $SCRIPTPATH
fi

echo "====== Coping templates"
cp -v -r $SCRIPTPATH/templates $dist_directory

echo "=== backend building finish"