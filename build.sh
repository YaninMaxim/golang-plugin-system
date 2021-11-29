#!/bin/bash

dir=$(pwd)
cd $dir/src/program && go build -o $dir/bin/main
chmod +x $dir/bin/main

for i in `find $dir/src/plugins -maxdepth 1 -type d -not -name "plugins"` 
do 
    cd $i
    go build -buildmode=plugin -o $dir/bin/plugins/$(pwd | xargs basename).so
done
