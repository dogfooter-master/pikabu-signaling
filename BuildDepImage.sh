#!/bin/sh
EXE=docker
unameOut="$(uname -s)"
if [ 0 -ne 0 ]; then
	EXE=${EXE}.exe
fi

${EXE} build --no-cache -t dermaster/pikabu-signaling:latest .
docker push dermaster/pikabu-signaling:latest
