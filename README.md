# cli-json-formatter

Small command to transform String into Json format.
This is a small project to test the Go language.

## Bookmarks

- https://go.dev/doc/
- https://github.com/spf13/cobra-cli/blob/main/README.md

## Init

// ...

```shell
export PATH="$(go env GOPATH)/bin:$PATH"
```

## Usage
### How to run a specific command

// ...

```sh
go build
```

// ...

```sh
./cli-json-formatter
```

> [!NOTE]
> The string argument to json should look like this : `'{"key": "value"}'`
> For the shell to correctly interpret the passed string, in order to format it in json

```sh
./cli-json-formatter format -u '{"key": "value"}'
```
