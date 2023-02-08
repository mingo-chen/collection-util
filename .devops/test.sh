#!/bin/bash

code_path="github.com/mingo-chen/collection-util"
go test -timeout 30s -count=2 -run ^Test ${code_path}/... -gcflags=all=-l -cover