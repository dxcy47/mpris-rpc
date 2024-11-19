#! /bin/bash

go build mpris-rpc.go
mkdir ~/.config/mpris-rpc
mv mpris-rpc.kdl ~/.config/mpris-rpc
mv mpris-rpc ~/.local/bin
