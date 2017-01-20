#!/bin/bash

go build -buildmode=plugin -o bin/myplug.so myplug
go build -o bin/myprog myprog
go build -o bin/hotloading hotloading
