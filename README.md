# xstream

![made with go](https://img.shields.io/badge/made%20with-Go-1E90FF.svg) [![go report card](https://goreportcard.com/badge/github.com/hueristiq/xstream)](https://goreportcard.com/report/github.com/hueristiq/xstream) [![release](https://img.shields.io/github/release/hueristiq/xstream?style=flat&color=1E90FF)](https://github.com/hueristiq/xstream/releases) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/xstream.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/xstream/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/xstream.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/xstream/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=1E90FF)](https://github.com/hueristiq/xstream/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-1E90FF.svg) [![contribution](https://img.shields.io/badge/contributions-welcome-1E90FF.svg)](https://github.com/hueristiq/xstream/blob/master/CONTRIBUTING.md)

`xstream` is a command-line interface (CLI) utility for handling data streams in Unix-like environments. It serves as a pivotal link between standard input and dual outputs - standard output and a file.

## Resources

* [Features](#features)
* [Installation](#installation)
	* [Install release binaries (Without Go Installed)](#install-release-binaries-without-go-installed)
	* [Install source (With Go Installed)](#install-source-with-go-installed)
		* [`go install ...`](#go-install)
		* [`go build ...` the development Version](#go-build--the-development-version)
* [Usage](#usage)
	* [Examples](#examples)
		* [Appending Unique Lines to File](#appending-unique-lines-to-file)
		* [Deduplicating Files](#deduplicating-files)
* [Contributing](#contributing)
* [Licensing](#licensing)
* [Credits](#credits)
    * [Contributors](#contributors)
    * [Similar Projects](#similar-projects)

## Features

* Writes incoming `stdin` to both `stdout` and file.
* Supports reading and writing to the same file.
* Supports appending desination.
* Supports deduplication.
* Cross-Platform (Windows, Linux & macOS).

## Installation

### Install release binaries (Without Go Installed)

Visit the [releases page](https://github.com/hueristiq/xstream/releases) and find the appropriate archive for your operating system and architecture. Download the archive from your browser or copy its URL and retrieve it with `wget` or `curl`:

* ...with `wget`:

	```bash
	wget https://github.com/hueristiq/xstream/releases/download/v<version>/xstream-<version>-linux-amd64.tar.gz
	```

* ...or, with `curl`:

	```bash
	curl -OL https://github.com/hueristiq/xstream/releases/download/v<version>/xstream-<version>-linux-amd64.tar.gz
	```

...then, extract the binary:

```bash
tar xf xstream-<version>-linux-amd64.tar.gz
```

> [!TIP]
> The above steps, download and extract, can be combined into a single step with this onliner
> 
> ```bash
> curl -sL https://github.com/hueristiq/xstream/releases/download/v<version>/xstream-<version>-linux-amd64.tar.gz | tar -xzv
> ```

> [!NOTE]
> On Windows systems, you should be able to double-click the zip archive to extract the `xstream` executable.

...move the `xstream` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

```bash
sudo mv xstream /usr/local/bin/
```

> [!NOTE]
> Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xstream` to their `PATH`.

### Install source (With Go Installed)

Before you install from source, you need to make sure that Go is installed on your system. You can install Go by following the official instructions for your operating system. For this, we will assume that Go is already installed.

#### `go install ...`

```bash
go install -v github.com/hueristiq/xstream/cmd/xstream@latest
```

#### `go build ...` the development Version

* Clone the repository

	```bash
	git clone https://github.com/hueristiq/xstream.git 
	```

* Build the utility

	```bash
	cd xstream/cmd/xstream && \
	go build .
	```

* Move the `xstream` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

	```bash
	sudo mv xstream /usr/local/bin/
	```

	Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xstream` to their `PATH`.


> [!CAUTION]
> While the development version is a good way to take a peek at `xstream`'s latest features before they get released, be aware that it may have bugs. Officially released versions will generally be more stable.

## Usage

To display help message for `xstream` use the `-h` flag:

```bash
xstream -h
```

help message:

```
          _
__  _____| |_ _ __ ___  __ _ _ __ ___
\ \/ / __| __| '__/ _ \/ _` | '_ ` _ \
 >  <\__ \ |_| | |  __/ (_| | | | | | |
/_/\_\___/\__|_|  \___|\__,_|_| |_| |_|
                                 v0.1.0

       with <3 by Hueristiq Open Source

USAGE:
 xstream [OPTIONS]

INPUT:
     --soak bool        soak up all input before writing to file

MANIPULATION:
     --trim bool        enable leading and trailing whitespace trimming

OUTPUT:
 -u, --unique bool      output unique lines
 -a, --append bool      append lines to output
 -q, --quiet bool       suppress output to stdout
 -p, --preview bool     preview new lines, without writing to file

pflag: help requested
```

### Examples

#### Appending Unique Lines to File

```
➜  cat stuff.txt
one
two
three

➜  cat new-stuff.txt
zero
one
two
three
four
five

➜  cat new-stuff.txt | xstream stuff.txt --append --unique
zero
four
five

➜  cat stuff.txt
one
two
three
zero
four
five

```

Note that the new lines added to `stuff.txt` are also sent to `stdout`, this allows for them to be redirected to another file:

```
➜  cat new-stuff.txt | xstream stuff.txt --append --unique > added-lines.txt
➜  cat added-lines.txt
zero
four
five
```

#### Deduplicating Files

```
➜  cat stuff.txt
zero
one
two
three
zero
four
five
five

➜  cat stuff.txt | xstream stuff.txt --soak --unique
zero
one
two
three
four
five

➜  cat stuff.txt
zero
one
two
three
four
five

```

Note the use of `--soak`, it makes the utility soak up all its input before writing to a file. This is useful for reading from and writing to the same file in a single pipeline.

## Contributing

[Issues](https://github.com/hueristiq/xstream/issues) and [Pull Requests](https://github.com/hueristiq/xstream/pulls) are welcome! **Check out the [contribution guidelines.](https://github.com/hueristiq/xstream/blob/master/CONTRIBUTING.md)**

## Licensing

This tool is distributed under the [MIT license](https://github.com/hueristiq/xstream/blob/master/LICENSE).

## Credits

### Contributors

Thanks to the amazing [contributors](https://github.com/hueristiq/xstream/graphs/contributors) for keeping this project alive.

[![contributors](https://contrib.rocks/image?repo=hueristiq/xstream&max=500)](https://github.com/hueristiq/xstream/graphs/contributors)

### Similar Projects

Thanks to similar open source projects - check them out, may fit in your workflow.

[anew](https://github.com/tomnomnom/anew) ◇ tee (coreutils) ◇ sponge (moreutils)