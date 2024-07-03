# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

```consolw
Below is the current list of plugins requiring updates for test.blog.ca.

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
```

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Creation of a variables file with the following values as per your environment:

```go
const (
server, path, site string = /* [target server], [wordpress path], [wordpress url] */
sender, recipient, user string = /* [mail sender], [mail recipient], [authorized user] */
)

// Allow only a predefined set of servers
var (
servers = []string{/* list of servers to test against */}
)
```

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```bash
go build -o [name] main.go
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o [name] main.go
```

## Run

```bash
./[program] [flag]
```

Currently the only working flags are **-p**, **-h**, and **-v**

Example:

```bash
./platypus -p
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/platypus/blob/main/LICENSE.md) and is part of the Public Domain.
