# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

``` zsh
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:22.9.0
wpackagist-plugin/tablepress:3.3
wpackagist-plugin/visualizer:4.0.1
wpackagist-plugin/woosidebars:1.4.6
wpackagist-theme/twentytwentyfive:1.4
```

## 📚 Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

Creation of a `test.json` file with the following values as per your environment:

``` json
{
    "address": "WordPress base url (no leading protocol)",
    "install": "Path on the server to the WordPress install",
    "recipient": "Email recipient(s) address(es)",
    "sender": "Address of Email sender ",
    "server": "Name of WordPress hosting server",
    "user": "User authorized to run the program"
}
```

## 🚧 Build

Before building the application, change the value of the `base` `temp`, and `jsons` constants to reflect your environment:

``` go
base, temp, jsons string = "/data/automation/", base + "temp/", base + "jsons/"
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
