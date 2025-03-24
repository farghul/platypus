# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

``` console
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
```

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Creation of a `env.json` file with the following values as per your environment:

``` json
{
    "blog": "Path on the server to the WordPress install",
    "recipient": "Email recipient(s) address(es)",
    "sender": "Email sender address",
    "server": "Server hosting WordPress",
    "site": "WordPress base url (no leading protocol)",
    "user": "User authorized to run the program"
}
```

## Build

Before building the application, change the value of the `base` constant to reflect your environment:

``` go
base string = "/data/automation/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac:

``` console
go build -o [name] .
```

### Linux:

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Options

``` console
-h, --help       Help Information
-r, --run        Run Program
-v, --version    Display App Version
```

## Run

``` console
./[program] [flag]
```

Example:

``` console
./platypus -r
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/platypus/blob/main/LICENSE.md) and is part of the Public Domain.
