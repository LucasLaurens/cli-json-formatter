# cli-json-formatter

Small command to transform String into Json format.
This is a small project to test the Go language.

<br/>

> [!NOTE]
> It's really just a tiny application for transforming a string into a json format, so you can learn more about the Go language.

<br/>

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

You need to compile the source of you program before being able to run the command.

```sh
go build
```


The next step is to check the help message because you're running the compile source command. It gives you all the child commands you've created and the options they contain.
```sh
./cli-json-formatter
```

> [!NOTE]
> The string argument to json should look like this : `'{"key": "value"}'`
> For the shell to correctly interpret the passed string, in order to format it in json

```sh
./cli-json-formatter format -u '{"key": "value"}'
```
