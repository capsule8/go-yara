# Repackaged go-yara

This package is a repackaged version of
https://github.com/hillu/go-yara which focuses on making it easier to
use.

Instead of requiring users to install the yara library on the build
system, this package already includes the library within it.  This
makes it very easy to use - simply add the dependency and build
normally.

The resulting binary will also be statically built and have no
external shared object dependencies. No need to mess with build flags:

```
~/go/src/github.com/Velocidex/go-yara$ go build _examples/simple-yara/simple-yara.go
~/go/src/github.com/Velocidex/go-yara$ ldd simple-yara
        linux-vdso.so.1 (0x00007ffd4085a000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fcd9867c000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fcd9828b000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fcd9889b000)
```

To build for windows, there is no need to build the yara library
first, just build your project with the usual steps:
```
~/go/src/github.com/Velocidex/go-yara$ GOOS=windows GOARCH=amd64 \
      CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
      go build _examples/simple-yara/simple-yara.go
```

![Logo](/goyara-logo.png)

# go-yara

[![PkgGoDev](https://pkg.go.dev/badge/github.com/hillu/go-yara/v4)](https://pkg.go.dev/github.com/hillu/go-yara/v4)
[![Travis](https://travis-ci.org/hillu/go-yara.svg?branch=master)](https://travis-ci.org/hillu/go-yara)
[![Go Report Card](https://goreportcard.com/badge/github.com/hillu/go-yara)](https://goreportcard.com/report/github.com/hillu/go-yara)

Go bindings for [YARA](https://virustotal.github.io/yara/), staying as
close as sensible to the library's C-API while taking inspiration from
the `yara-python` implementation.

## Build/Installation

On Unix-like systems, _libyara_ version 4, corresponding header files,
and _pkg-config_ must be installed. Adding _go-yara_ v4 to a project
with Go Modules enabled, simply add the proper dependency…

``` go
import "github.com/hillu/go-yara/v4"
```

…and rebuild your package.

If _libyara_ has been installed to a custom location, the
`PKG_CONFIG_PATH` environment variable can be used to point
_pkg-config_ at the right `yara.pc` file.

For anything more complicated, refer to the "Build Tags" section
below. Instructions for cross-building _go-yara_ for different
operating systems or architectures can be found in
[README.cross-building.md](README.cross-building.md).

To build _go-yara_ on Windows, a GCC-based build environment is
required, preferably one that includes _pkg-config_. The 32-bit and
64-bit MinGW environments provided by the [MSYS2](https://msys2.org/)
provide such an environment.

## Build Tags

### Static builds

The build tag `yara_static` can be used to tell the Go toolchain to
run _pkg-config_ with the `--static` switch. This is not enough for a
static build; the appropriate linker flags (e.g. `-extldflags
"-static"`) still need to be passed to the _go_ tool.

### Building without _pkg-config_

The build tag `yara_no_pkg_config` can be used to tell the Go toolchain not
to use _pkg-config_'s output. In this case, any compiler or linker
flags have to be set via the `CGO_CFLAGS` and `CGO_LDFLAGS`
environment variables, e.g.:

```
export CGO_CFLAGS="-I${YARA_SRC}/libyara/include"
export CGO_LDFLAGS="-L${YARA_SRC}/libyara/.libs -lyara"
go install -tags yara_no_pkg_config github.com/hillu/go-yara
```

## YARA 4.x vs. earlier versions

This version of _go-yara_ can only be used with YARA 4.0 or later.

Versions of _go-yara_ compatible with YARA 3.11 are available via the
`v3.x` branch or tagged `v3.*` releases.

Versions of _go-yara_ compatible with earlier 3.x versions of YARA are
available via the `v1.x` branch or tagged `v1.*` releases.

## License

BSD 2-clause, see LICENSE file in the source distribution.

## Author

Hilko Bengen <<bengen@hilluzination.de>>
