# xt

[![release](https://img.shields.io/github/release/hueristiq/xt?style=flat&color=0040ff)](https://github.com/hueristiq/xt/releases) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/hueristiq/xt/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/xt.svg?style=flat&color=0040ff)](https://github.com/hueristiq/xt/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/xt.svg?style=flat&color=0040ff)](https://github.com/hueristiq/xt/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-0040ff.svg)](https://github.com/hueristiq/xt/blob/master/CONTRIBUTING.md)

`xt` is a CLI based container management tool to efficiently build and interact with containers on docker environments.

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

* [ ] Copy standard input to each FILE, and also to standard output.

## Installation

### Install release binaries (Without Go Installed)

Visit the [releases page](https://github.com/hueristiq/xt/releases) and find the appropriate archive for your operating system and architecture. Download the archive from your browser or copy its URL and retrieve it with `wget` or `curl`:

* ...with `wget`:

	```bash
	wget https://github.com/hueristiq/xt/releases/download/v<version>/xt-<version>-linux-amd64.tar.gz
	```

* ...or, with `curl`:

	```bash
	curl -OL https://github.com/hueristiq/xt/releases/download/v<version>/xt-<version>-linux-amd64.tar.gz
	```

...then, extract the binary:

```bash
tar xf xt-<version>-linux-amd64.tar.gz
```

> **TIP:** The above steps, download and extract, can be combined into a single step with this onliner
> 
> ```bash
> curl -sL https://github.com/hueristiq/xt/releases/download/v<version>/xt-<version>-linux-amd64.tar.gz | tar -xzv
> ```

**NOTE:** On Windows systems, you should be able to double-click the zip archive to extract the `xt` executable.

...move the `xt` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

```bash
sudo mv xt /usr/local/bin/
```

**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xt` to their `PATH`.

### Install source (With Go Installed)

Before you install from source, you need to make sure that Go is installed on your system. You can install Go by following the official instructions for your operating system. For this, we will assume that Go is already installed.

#### `go install ...`

```bash
go install -v github.com/hueristiq/xt/cmd/xt@latest
```

#### `go build ...` the development Version

* Clone the repository

	```bash
	git clone https://github.com/hueristiq/xt.git 
	```

* Build the utility

	```bash
	cd xt/cmd/xt && \
	go build .
	```

* Move the `xt` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

	```bash
	sudo mv xt /usr/local/bin/
	```

	**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xt` to their `PATH`.


**NOTE:** While the development version is a good way to take a peek at `xt`'s latest features before they get released, be aware that it may have bugs. Officially released versions will generally be more stable.

## Usage

To display help message for `xt` use the `-h` flag:

```bash
xt -h
```

help message:

```
       _   
__  _ | |_ 
\ \/ /  __|
 >  < | |_ 
/_/\_\ \__| v0.0.0

USAGE:
  xt [OPTIONS]

OPTIONS:
 -q, --quite bool                 quiet mode (no output at all)
     --dry-run bool               don't append anything to the file, just print the new lines to stdout
     --trim bool                  trim leading and trailing whitespace before comparison

```

## Contributing

[Issues](https://github.com/hueristiq/xt/issues) and [Pull Requests](https://github.com/hueristiq/xt/pulls) are welcome! **Check out the [contribution guidelines.](./CONTRIBUTING.md)**

## Licensing

This tool is distributed under the [MIT license](https://github.com/hueristiq/xt/blob/master/LICENSE).