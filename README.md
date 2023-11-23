# xbridge

![made with go](https://img.shields.io/badge/made%20with-Go-1E90FF.svg) [![go report card](https://goreportcard.com/badge/github.com/hueristiq/xbridge)](https://goreportcard.com/report/github.com/hueristiq/xbridge) [![release](https://img.shields.io/github/release/hueristiq/xbridge?style=flat&color=1E90FF)](https://github.com/hueristiq/xbridge/releases) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/xbridge.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/xbridge/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/xbridge.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/xbridge/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=1E90FF)](https://github.com/hueristiq/xbridge/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-1E90FF.svg) [![contribution](https://img.shields.io/badge/contributions-welcome-1E90FF.svg)](https://github.com/hueristiq/xbridge/blob/master/CONTRIBUTING.md)

`xbridge` is a command-line interface (CLI) utility for handling data streams. It serves as a pivotal link between standard input and dual outputs - standard output and a file.

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
* [Credits](#credits)
    * [Contributors](#contributors)
    * [Similar Projects](#similar-projects)

## Features

* Dual outputs - standard output and a file.
* Cross-Platform (Windows, Linux & macOS).

## Installation

### Install release binaries (Without Go Installed)

Visit the [releases page](https://github.com/hueristiq/xbridge/releases) and find the appropriate archive for your operating system and architecture. Download the archive from your browser or copy its URL and retrieve it with `wget` or `curl`:

* ...with `wget`:

	```bash
	wget https://github.com/hueristiq/xbridge/releases/download/v<version>/xbridge-<version>-linux-amd64.tar.gz
	```

* ...or, with `curl`:

	```bash
	curl -OL https://github.com/hueristiq/xbridge/releases/download/v<version>/xbridge-<version>-linux-amd64.tar.gz
	```

...then, extract the binary:

```bash
tar xf xbridge-<version>-linux-amd64.tar.gz
```

> **TIP:** The above steps, download and extract, can be combined into a single step with this onliner
> 
> ```bash
> curl -sL https://github.com/hueristiq/xbridge/releases/download/v<version>/xbridge-<version>-linux-amd64.tar.gz | tar -xzv
> ```

**NOTE:** On Windows systems, you should be able to double-click the zip archive to extract the `xbridge` executable.

...move the `xbridge` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

```bash
sudo mv xbridge /usr/local/bin/
```

**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xbridge` to their `PATH`.

### Install source (With Go Installed)

Before you install from source, you need to make sure that Go is installed on your system. You can install Go by following the official instructions for your operating system. For this, we will assume that Go is already installed.

#### `go install ...`

```bash
go install -v github.com/hueristiq/xbridge/cmd/xbridge@latest
```

#### `go build ...` the development Version

* Clone the repository

	```bash
	git clone https://github.com/hueristiq/xbridge.git 
	```

* Build the utility

	```bash
	cd xbridge/cmd/xbridge && \
	go build .
	```

* Move the `xbridge` binary to somewhere in your `PATH`. For example, on GNU/Linux and OS X systems:

	```bash
	sudo mv xbridge /usr/local/bin/
	```

	**NOTE:** Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `xbridge` to their `PATH`.


**NOTE:** While the development version is a good way to take a peek at `xbridge`'s latest features before they get released, be aware that it may have bugs. Officially released versions will generally be more stable.

## Usage

To display help message for `xbridge` use the `-h` flag:

```bash
xbridge -h
```

help message:

```

      _          _     _
__  _| |__  _ __(_) __| | __ _  ___
\ \/ / '_ \| '__| |/ _` |/ _` |/ _ \
 >  <| |_) | |  | | (_| | (_| |  __/
/_/\_\_.__/|_|  |_|\__,_|\__, |\___|
                         |___/
                              v0.0.0

    with <3 by Hueristiq Open Source

USAGE:
 xbridge [OPTIONS]

MANIPULATION:
     --trim bool        enable leading and trailing whitespace trimming

OUTPUT:
 -p, --preview bool     enable preview mode i.e show new lines, without writing the changes
 -s, --silent bool      enable silent mode i.e suppress stdout output

pflag: help requested
```

## Contributing

[Issues](https://github.com/hueristiq/xbridge/issues) and [Pull Requests](https://github.com/hueristiq/xbridge/pulls) are welcome! **Check out the [contribution guidelines.](https://github.com/hueristiq/xbridge/blob/master/CONTRIBUTING.md)**

## Licensing

This tool is distributed under the [MIT license](https://github.com/hueristiq/xbridge/blob/master/LICENSE).

## Credits

### Contributors

Thanks to the amazing [contributors](https://github.com/hueristiq/xbridge/graphs/contributors) for keeping this project alive.

[![contributors](https://contrib.rocks/image?repo=hueristiq/xbridge&max=500)](https://github.com/hueristiq/xbridge/graphs/contributors)

### Similar Projects

Thanks to similar open source projects - check them out, may fit in your workflow.

[anew](https://github.com/tomnomnom/anew) ◇ tee (coreutils) ◇ sponge (moreutils)