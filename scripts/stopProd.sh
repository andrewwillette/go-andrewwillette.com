#!/usr/bin/env bash

# stops the background go server
kill $(ps|grep willette_api|awk '! /grep/ {print $1}')

# stops the background express server
# fix this for ubuntu deployment
kill $(ps|grep node|awk '{print $1}')
