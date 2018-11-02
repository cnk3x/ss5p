#!/bin/sh
if [ -n "${ip}" ]; then
    /ss5p --usr "${usr}" --pwd "${pwd}" --port=8080 --ip=${ip}
else
    /ss5p --usr "${usr}" --pwd "${pwd}" --port=8080
fi