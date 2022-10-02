#!/bin/bash
DIR=$(cd $(dirname $0) && pwd)
echo "DIR $DIR"
cd $DIR
./bac_gui

#read -p "Press [Enter] key to start backup..."