#!/bin/bash

echo "Performance testing - 100 executions:"
time (for i in $(seq 100); do python ./run; done) > ../eawasy-go.run 2>&1
cat ../eawasy-go.run
/usr/bin/time -v ./run
