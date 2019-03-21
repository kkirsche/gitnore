# gitnore

Are you tired of choosing overly bare .gitignore files from Github's interface? So was I. That's why I made gitnore, a tool to list, view, and download gitignore files from https://github.com/github/gitignore and merge the contents of the files you choose into your gitignore, allowing you to build full featured, rich .gitignore files quickly and easily.

## Installation

### From Source

```
$ go get -u https://github.com/kkirsche/gitnore
```

## Usage

### Initialization

If you want to provide your personal access token via the CLI, you can initialize your configuration using:

```
$ gitnore init -t="example-personal-access-token"
```

This though has some security concerns if on a shared machine, and as such you can instead use the following command to be prompted:

```
$ gitnore init
```

### Updating Your Token

You can either manually edit your configuration file, or use the `writeConfig` command.

**WARNING: This will overwrite your configuration file!**

```
$ gitnore writeConfig
```

## List Available Gitignore Files

```
$ gitnore list
```

## Preview Gitignore File Content

```
$ gitnore preview gitignore-file-name
```

for example:

```
$ gitnore preview Go
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
```

## Add a Gitignore File's Contents to Your Local Gitignore File

```
$ gitnore add gitignore-file-name
```
