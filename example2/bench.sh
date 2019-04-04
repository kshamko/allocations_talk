#!/bin/sh

go test -benchmem -run=none -bench Assignment -memprofile mem.out