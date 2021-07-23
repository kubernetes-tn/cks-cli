
#!/usr/bin/make -f
include .env
# auto populate env vars from .env file
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' .env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"
MKDOCS_IMAGE := ktn/mkdocs-material:dev
MKDOCS_PORT := 8000

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o cks .

.PHONY: build-all
build-all:
	sh .build-all

release-binary:
	make build-all
	aws s3 sync ./bin/ s3://$(BUCKET_BINARIES)/cks-cli/$(VERSION) --delete --region eu-west-1 
	aws s3 sync s3://$(BUCKET_BINARIES)/cks-cli/$(VERSION) s3://$(BUCKET_BINARIES)/cks-cli/latest --delete --region eu-west-1 
mkdocs-serve:
	docker build -t $(MKDOCS_IMAGE) -f docs/build/Dockerfile docs/build
	docker run --name mkdocs-serve --rm --env-file .env -v $(PWD):/docs -p $(MKDOCS_PORT):8000 $(MKDOCS_IMAGE)

mkdocs-gen:
	docker build -t $(MKDOCS_IMAGE) -f docs/build/Dockerfile docs/build
	docker run --name mkdocs-gen --rm --env-file .env -v $(PWD):/docs --entrypoint=mkdocs $(MKDOCS_IMAGE) build --site-dir public

release-docs:
	make mkdocs-gen
	aws s3 sync ./public/ s3://$(shell bash deploy/cloudformation/get_bucket_website_name.sh)/ --delete --region ap-southeast-1

release-infra:
	bash deploy/cloudformation/apply.sh
release:
	make release-binary
	make release-docs


# https://stackoverflow.com/a/41629516/747579
