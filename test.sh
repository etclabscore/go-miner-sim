#!/usr/bin/env bash

mkdir -p test_results
go test -v -count=1 -run TestPlotting . |& tee test_results/log_$(date +%s).txt
