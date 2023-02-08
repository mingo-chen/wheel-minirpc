#!/bin/bash

code_path="github.com/mingo-chen/wheel-minirpc"
go test -timeout 30s -count=2 -run ^Test ${code_path}/... -gcflags=all=-l -cover