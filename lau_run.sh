#!/bin/sh

# stop running
./lau_stop.sh

# start running
nohup ./lau >>lau.log 2>&1 &