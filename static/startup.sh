#!/usr/bin/env bash

mypath=$(dirname $0)
cd ${mypath}

if [ $# -eq 0 ]; then
    echo 'Illegal input, you can choose:'
    echo '  [--help] help document'
    echo '  [--server] launch a server'
    echo '  [--executor] launch a executor'
    echo '  [--client] launch a client'
    exit
fi

arg0=$1

if [ "$arg0" = "--executor" ]; then
    cd executor
    ./executor
fi

if [ "$arg0" = "--server" ]; then
    cd server
    ./server
fi

if [ "$arg0" = "--client" ]; then
    cd client
    ./client
fi

if [ "$arg0" = "--help" ]; then
    echo 'no help!'
fi

