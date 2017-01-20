#!/bin/bash

docker run -it --rm -p 8080:8080 -v $PWD/..:/go go-plugins-talk
