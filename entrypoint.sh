#!/bin/sh
set -e

if [ "$1" = 'akita' ]; then
        /akita apidump \
                --service brill-wtf \
                --filter "port 3333" \
                -u root \
                -c ./rest
fi