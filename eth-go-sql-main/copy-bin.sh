#!/bin/bash

cp -f ./build/producer ../build/eth/
cp -f ./build/producer ../build/bsc/
cp -f ./build/producer ../build/matic/
cp -f ./build/producer ../build/moonriver/

cp -f ./build/consumer ../build/eth/
cp -f ./build/consumer ../build/bsc/
cp -f ./build/consumer ../build/matic/
cp -f ./build/consumer ../build/moonriver/

cp -f ./build/repair ../build/eth/
cp -f ./build/repair ../build/bsc/
cp -f ./build/repair ../build/matic/
cp -f ./build/repair ../build/moonriver/

cp -f ./build/check ../build/eth/
cp -f ./build/check ../build/bsc/
cp -f ./build/check ../build/matic/
cp -f ./build/check ../build/moonriver/
