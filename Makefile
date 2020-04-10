GIT_TAG?= $(shell git describe --always --tags)
BIN = badbot
EXE = badbot.exe
FMT_CMD = $(gofmt -s -l -w $(find . -type f -name '*.go' -not -path './vendor/*') | tee /dev/stderr)
IMAGE_REPO = badbot
BUILDFLAGS := '-w -s'
CGO_ENABLED = 0
GO := GO111MODULE=on go
GO_NOMOD :=GO111MODULE=off go
BUILD_DATE := $(shell date --rfc-3339=seconds | sed 's/ /T/')
LDFLAGS := '-w -s -X github.com/znk3r/badbot/cmd.Version=Dev -X github.com/znk3r/badbot/cmd.GitTag=$(GIT_TAG) -X github.com/znk3r/badbot/cmd.BuildDate=$(BUILD_DATE)'
COVER_OUT = coverage.out
COVER_HTML = coverage.html

OBJ_COLOR = \033[0;37m\033[1m
INFO_COLOR = \033[0;96m
OK_COLOR = \033[0;92m
ERROR_COLOR = \033[0;91m
WARN_COLOR = \033[0;93m
NO_COLOR = \033[m\033[21m

default:
	$(MAKE) build

lint: 
	@echo -e "$(OBJ_COLOR)[lint] $(INFO_COLOR)Linting$(NO_COLOR)"
	golint ./...
	@echo -e "$(OBJ_COLOR)[lint] $(INFO_COLOR)Vetting$(NO_COLOR)"
	@$(GO) vet ./...

test:
	$(MAKE) test-unit

test-unit:
	richgo test -v ./...

test-coverage:
	@echo -e "$(OBJ_COLOR)[coverage] $(INFO_COLOR)Generating code coverage$(NO_COLOR)"
	go test -race -coverprofile=$(COVER_OUT) -covermode=atomic ./...
	@echo -e "$(OBJ_COLOR)[coverage] $(INFO_COLOR)Generating html coverage$(NO_COLOR)"
	go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)

build:
	go build -ldflags $(LDFLAGS)

build-linux:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)

build-win:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(EXE)

clean:
	rm -rf build vendor dist
	rm -f release image $(COVER_OUT) $(COVER_HTML) $(BIN) $(EXE)

release: 
	@echo -e "$(OBJ_COLOR)[release] $(INFO_COLOR)Releasing badbot binary $(GIT_TAG)$(NO_COLOR)"
	goreleaser release
