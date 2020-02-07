# taffmat

Go package for reading and writing Teac TAFFmat files.

[![GoDoc][godoc image]][godoc link]
[![Go Report Card][report badge]][report card]
[![Build Status][travis image]][travis link]
[![Coverage Status][coveralls image]][coveralls link]
[![License Badge][license image]][LICENSE.txt]

## About the TAFFmat file format

TAFFmat is Teac's proprietary file format used to store data from their
LX series and other data recorders.

According to the Teac "LX Series Recording Unit Instruction Manual":

>  TAFFmat (an acronym for Teac Data Acquisition File Format) is a
>  file format composed of the following:
>
>  * a data file containing A/D (analog to digital) converted data. The
>    file is binary format with the extension dat.
>  * a header file containing information such as recording
>    conditions. The file is in text format with the extension hdr.

TAFFmat is a trademark of Teac Corporation.

### Data Recorders Using TAFFmat

The following data recorders store their data in the TAFFmat file format:

* Teac [LX-10/20][]
* Teac [LX-110/120][]
* Teac [WX-7000 Series][]
* Teac [es8][]

# Installation

```bash
$ go get github.com/gotmc/taffmat
```

# Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/taffmat>
- <http://localhost:6060/pkg/github.com/gotmc/taffmat/> after running `$
  godoc -http=:6060`

# Contributing

To contribute, fork [taffmat][], create a feature branch, and then submit a
[pull request][].

# Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[taffmat][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[coveralls image]: http://img.shields.io/coveralls/gotmc/taffmat/master.svg
[coveralls link]: https://coveralls.io/r/gotmc/taffmat
[es8]: http://teac-ipd.com/data-recorders/es8/
[godoc image]: https://godoc.org/github.com/gotmc/libusb?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/libusb
[taffmat]: https://github.com/gotmc/taffmat
[LICENSE.txt]: https://github.com/gotmc/taffmatb/blob/master/LICENSE.txt
[license image]: https://img.shields.io/badge/license-MIT-blue.svg
[LX-10/20]: http://www.teac.co.jp/en/industry/measurement/datarecorder/lx10/index.html
[LX-110/120]: http://teac-ipd.com/data-recorders/lx-110120/
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/gotmc/taffmat
[report card]: https://goreportcard.com/report/github.com/gotmc/taffmat
[travis image]: http://img.shields.io/travis/gotmc/taffmat/master.svg
[travis link]: https://travis-ci.org/gotmc/taffmat
[WX-7000 Series]: http://teac-ipd.com/wx-7000/
