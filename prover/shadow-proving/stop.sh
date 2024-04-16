#!/bin/bash

# pkill -9 shadow-proving
ps -ef | grep 'shadow-proving' | grep -v grep | awk '{print $2}' | xargs -r kill -9
