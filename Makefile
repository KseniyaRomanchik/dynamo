export IMAGE_NAME=golang:1.13.5-alpine3.11
export WORKDIR=/go/src/dynamo
export COMMIT=$(shell git rev-parse HEAD)
export DATE=$(shell date +%d-%m-%Y__%T)

build:
	@echo '*** BUILD ***'
	@sudo docker run --env CGO_ENABLED=0 --name dynamo-build --rm -v $(PWD)\:$(WORKDIR) $(IMAGE_NAME) /bin/sh \
	-c 'cd $(WORKDIR) && mkdir -p cmd && \
	 go build -ldflags \
	 "-X dynamo/flags.Image=$(IMAGE_NAME) \
	 -X dynamo/flags.Commit=$(COMMIT) \
	 -X dynamo/flags.Time=$(DATE)" \
	 -o ./cmd'

install: build
	@bash -c "echo '*** delete ~/bin/dynamo' ; rm -fr ~/bin/dynamo "
	@bash -c "echo '*** cp new dynamo ';  cp cmd/dynamo ~/bin/dynamo"

install_global: build
	@bash -c "echo '*** delete /usr/bin/dynamo' ; rm -fr /usr/bin/dynamo "
	@bash -c "echo '*** cp new dynamo ';  cp cmd/dynamo /usr/bin/dynamo"