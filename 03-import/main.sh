#!/usr/bin/env bash
echo -n "The command line arguments are: ${0}"
for i in "${@}"; do
    echo -n " ${i}"
done
echo
echo "The SHELL is: ${SHELL}"
