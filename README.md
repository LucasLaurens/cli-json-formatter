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

To be able to use the Cobra package executable, you need to follow the documentation provided and then remember to export the path. It adds the Go workspace's bin directory to your system's PATH environment variable.

<br/>

> [!NOTE]
> This allows you to run Go installed binaries (e.g., tools installed with go install) from anywhere in your terminal without needing to specify their full path.

<br/>


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

<br/>

> [!NOTE]
> The string argument to json should look like this : `'{"key": "value"}'`
> For the shell to correctly interpret the passed string, in order to format it in json

<br/>

```sh
./cli-json-formatter format -u '{"key": "value"}'
```
