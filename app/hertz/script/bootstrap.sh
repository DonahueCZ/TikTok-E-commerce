#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=order
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}