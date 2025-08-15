# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

``` zsh
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
```

## 📚 Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

Creation of a `test.json` file with the following values as per your environment:

``` json
{
    "address": "WordPress base url (no leading protocol)",
    "install": "Path on the server to the WordPress install",
    "recipient": "Email recipient(s) address(es)",
    "sender": "Email sender address",
    "server": "Server hosting WordPress",
    "user": "User authorized to run the program"
}
```

## 🚧 Build

Before building the application, change the value of the `base` constant to reflect your environment:

``` go
base string = "/data/automation/"
```

And the value of the `repo` constant to point to a location for the config json files:

``` go
repo string = base + "bitbucket/desso-automation-conf/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac:

``` zsh
go build -o [name] .
```

### Linux:

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## 🏃 Run

``` zsh
./platypus -r
```

## 🎏 Available Flags

| Command               | Action                      |
|:----------------------|:----------------------------|
|    `-h, --help`       |   Help information          |
|    `-r, --run`        |   Run program               |
|    `-v, --version`    |   Display program version   |

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/farghul/platypus/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
