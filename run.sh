#!/bin/sh

SRC_FOLDER="cmd"
BIN_FOLDER="bin"
BIN_NAME="iot-electronic-wallet"

mkdir -p "$BIN_FOLDER"
mkdir -p "$SRC_FOLDER"
go build -o $BIN_FOLDER/$BIN_NAME $SRC_FOLDER/*
./$BIN_FOLDER/$BIN_NAME