#!/bin/bash

mkdir tmp
cp -R build/ tmp/build
cp app.yaml tmp/
cp .gcloudignore tmp/.gcloudignore
cp .gitignore tmp/.gitignore

cd tmp/
gcloud app deploy

cd ../

rm -rf tmp