#!/usr/bin/env bash

# stops the background go server
kill $(ps|grep go|awk '/var/{print $1}')
