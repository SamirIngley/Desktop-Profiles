#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d=pkg -pv=$(VERSION)

release:
	ghr  -u –desc=An app and url launcher  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
