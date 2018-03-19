# go cmd
> a go sub-comand package.

## example
see example director.
```
$ cd example
$ go build
$ ./example  # or ./example --help or ./example help or ./example -h or ./example h
Usage of example:

        example command [arguments]

The commands are:

        print           print some text
        version         print example version

Use "example help [command]" for more information about a command.
```
```bash
$ ./example abc
example: unknown subcommand "abc"
Run 'example help' for usage.
```

```bash
$ ./example print -c # an argument after "-c" is required.
flag needs an argument: -c
print some text, e.g. Hello World.

Useage of command "print":

        example print [options]

Options:

  -c string
        text for print. (default "default text")
default text
```

```bash
$ ./example  print
default text
$ ./example  print -c hello
hello
```
```bash
$ ./example print -c hello
hello
```

```bash
$ ./example version
version  0.1.0.
Author   abc@example.com
```