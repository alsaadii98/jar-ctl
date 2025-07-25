APP_NAME=jar-ctl

.PHONY: all build run clean test

all: build

build:
	go build -o jar-ctl main.go

run:
	./jar-ctl $(filter-out $@,$(MAKECMDGOALS))

clean:
	rm -f jar-ctl
