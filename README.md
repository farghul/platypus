# Platypus

Platypus is a WordPress plugin update search tool. It runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

![Platypus](platypus.webp)

```consolw
Below is the current list of plugins requiring updates for test.blog.ca

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
```

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Creation of a separate variables file with the following values as per your environment:

```go
// WordPress installation specific values 
const (
	server    string = /* [Server hosting WordPress] */
	path      string = /* [Path on the server to the WordPress install] */
	site      string = /* [WordPress base url (no leading protocol)] */
	sender    string = /* [email sender address] */
	recipient string = /* [email recipient(s) address(es)] */
	user      string = /* [user authorized to run the program] */
)

// Predefined list of servers
var (
	servers = []string{/* list of servers to test against */}
)
```

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```bash
go build -o [name] .
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o [name] .
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
