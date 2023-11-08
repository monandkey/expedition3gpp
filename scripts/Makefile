BINNAME:=e3gppcli
BINDIR:=../build/bin
GOARCH:=amd64
GO_FILES:=../apps/cli/main.go

.PHONY: build_gui build_cli

build_gui:
	wails build -platform windows/amd64

build_cli:
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME).exe $(GO_FILES)
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME) $(GO_FILES)
