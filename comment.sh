#!/bin/bash

# Quickly remove a word from the dictionary
if [ -z "$1" ]; then
    echo "Usage: $0 [word]"
fi

sed -i "/^$1$/I s/^#*/#/" TWL06.txt
