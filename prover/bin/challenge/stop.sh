#!/bin/bash

# pkill -9 challenge-handler
ps -ef | grep 'challenge-handler' | grep -v grep | awk '{print $2}' | xargs -r kill -9
