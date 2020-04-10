# badbot
Discord bot to add music to online table rpg

## Dev

### How to build?

```bash
$ make
```

You can also build a linux amd64 compatible version (GOOS=linux GOARCH=amd64) from any host system with:

```bash
$ make build-linux
```

Or a Windows compatible one with:

```bash
$ make build-win
```

### How to test?

To run tests, you need to install and have in your path some go packages:

```bash
$ go get -u golang.org/x/lint/golint
$ go get -u github.com/kyoh86/richgo
```

### How to release to github

First, make sure you have installed the [goreleaser](https://github.com/goreleaser/goreleaser) tool.

After tagging, you can create the GitHub release with the different binaries attached. To do this, you need to create your own GitHub Enterprise token and export it as an environment variable.

```bash
$ git tag -a v1.4.1 -m "< RELEASE MESSAGE >"
$ export GITHUB_TOKEN=<YOUR GITHUB TOKEN>
$ make release
```

