# xstreamin

[![release](https://img.shields.io/github/release/hueristiq/xstreamin?style=flat&color=0040ff)](https://github.com/hueristiq/xstreamin/releases) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/hueristiq/xstreamin/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/xstreamin.svg?style=flat&color=0040ff)](https://github.com/hueristiq/xstreamin/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/xstreamin.svg?style=flat&color=0040ff)](https://github.com/hueristiq/xstreamin/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-0040ff.svg)](https://github.com/hueristiq/xstreamin/blob/master/CONTRIBUTING.md)

`xstreamin` is a CLI utility to process and append unique lines of input to a specified file.

## Resources

* [Features](#features)
* [Installation](#installation)
	* [Install release binaries (Without Go Installed)](#install-release-binaries-without-go-installed)
	* [Install source (With Go Installed)](#install-source-with-go-installed)
		* [`go install ...`](#go-install)
		* [`go build ...` the development Version](#go-build--the-development-version)
* [Usage](#usage)
* [Contributing](#contributing)
* [Licensing](#licensing)

## Features

* Whitespace trimming to ensures consistent formatting by removing leading and trailing spaces from lines.
* Dry-Run capability to allow users to preview the lines that would be appended, without making any changes to the file.
* Efficient Duplication Check to filter out lines that already exist in the target file, ensuring no duplicates.
* Provides an option for quite/silent operation, suppressing output lines for a cleaner console experience.
* Cross-Platform (Windows, Linux & macOS).

## Installation

### Install release binaries (Without Go Installed)

Visit the [releases page](https://github.com/hueristiq/xstreamin/releases) and find the appropriate archive for your operating system and architecture. Download the archive from your browser or copy its URL and retrieve it with `wget` or `curl`:

* ...with `wget`:

	```bash
	wget https://github.com/hueristiq/xstreamin/releases/download/v<version>/xstreamin-<version>-linux-amd64.tar.gz
	```

* ...or, with `curl`:

	```bash
	curl -OL https://github.com/hueristiq/xstreamin/releases/download/v<version>/xstreamin-<version>-linux-amd64.tar.gz
	```

...then, extract the binary:

```bash
tar xf xstreamin-<version>-linux-amd64.tar.gz
```

> **TIP:** The above steps, download and extract, can be combined into a single step with this onliner
> 
> ```bash
> curl -sL https://github.com/hueristiq/xstreamin/releases/download/v<version>/xstreamin-<version>-linux-amd64.tar.gz | tar -xzv
> ```

**NOTE:** On Windows systems, you should be able to double-click the zip archive to extract the `xstreamin` executable.

...move the `xstreamin` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

```bash
sudo mv xstreamin /usr/local/bin/
```

**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xstreamin` to their `PATH`.

### Install source (With Go Installed)

Before you install from source, you need to make sure that Go is installed on your system. You can install Go by following the official instructions for your operating system. For this, we will assume that Go is already installed.

#### `go install ...`

```bash
go install -v github.com/hueristiq/xstreamin/cmd/xstreamin@latest
```

#### `go build ...` the development Version

* Clone the repository

	```bash
	git clone https://github.com/hueristiq/xstreamin.git 
	```

* Build the utility

	```bash
	cd xstreamin/cmd/xstreamin && \
	go build .
	```

* Move the `xstreamin` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

	```bash
	sudo mv xstreamin /usr/local/bin/
	```

	**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xstreamin` to their `PATH`.


**NOTE:** While the development version is a good way to take a peek at `xstreamin`'s latest features before they get released, be aware that it may have bugs. Officially released versions will generally be more stable.

## Usage

To display help message for `xstreamin` use the `-h` flag:

```bash
xstreamin -h
```

help message:

```

          _                            _
__  _____| |_ _ __ ___  __ _ _ __ ___ (_)_ __
\ \/ / __| __| '__/ _ \/ _` | '_ ` _ \| | '_ \
 >  <\__ \ |_| | |  __/ (_| | | | | | | | | | |
/_/\_\___/\__|_|  \___|\__,_|_| |_| |_|_|_| |_|
                                         v0.0.0

               with <3 by Hueristiq Open Source

USAGE:
 xstreamin [OPTIONS]

OPTIONS:
 -d, --dry-run bool     preview lines that would be appended, without writing the changes
 -q, --quite bool       enable quiet mode i.e. suppress output lines
     --trim bool        trim leading and trailing whitespace before comparison

pflag: help requested
```

## Contributing

[Issues](https://github.com/hueristiq/xstreamin/issues) and [Pull Requests](https://github.com/hueristiq/xstreamin/pulls) are welcome! **Check out the [contribution guidelines.](https://github.com/hueristiq/xstreamin/blob/master/CONTRIBUTING.md)**

## Licensing

This tool is distributed under the [MIT license](https://github.com/hueristiq/xstreamin/blob/master/LICENSE).