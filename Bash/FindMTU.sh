#!/bin/bash

host="1.1.1.1"
if [ $# -gt 0 ]; then
    host=$1
fi
echo "Finding MTU max:"
echo "================"
min=68
max=9000

while [ $min -le $max ]; do
        mtu=$(((min+max)/2))
        ping -W 3 -c 1 -M do -s $mtu $host > /dev/null 2>&1
    if [ $? -eq 0 ]; then
        min=$((mtu+1))
        echo -n "+"
    else
        max=$((mtu-1))
        echo -n "-"
    fi
done
mtu=$((max+28))
echo -e "\nMaximum MTU value for $host: $mtu"