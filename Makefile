#!/usr/bin/make -f

test:
	go test -short $(ARGS) ./...

day:
	@cd day$(DAY); rm -f ./day$(DAY); go build && ./day$(DAY); rm -f ./day$(DAY)

all:
	for day in $$(seq -f "%02g" 1 25) ; do \
		echo $$day ; \
		make day DAY=$$day ; \
	done
