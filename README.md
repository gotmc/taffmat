# taffmat

Go package for reading and writing TEAC TAFFmat files.

[![GoDoc][godoc image]][godoc link]
[![Go Report Card][report badge]][report card]
[![Build Status][travis image]][travis link]
[![License Badge][license image]][LICENSE.txt]

## About the TAFFmat file format

TAFFmat is TEAC's proprietary file format used to store data from their
LX series and other data recorders.

According to the TEAC "LX Series Recording Unit Instruction Manual":

>  TAFFmat (an acronym for TEAC Data Acquisition File Format) is a
>  file format composed of the following:
>
>  - a data file containing A/D (analog to digital) converted data. The
>    file is binary format with the extension dat.
>  - a header file containing information such as recording
>    conditions. The file is in text format with the extension hdr.

TAFFmat is a trademark of [TEAC Corporation][teac].

### Data Recorders Using TAFFmat

The following data recorders store their data in the TAFFmat file format:

- TEAC [LX-10/20][]
- TEAC [LX-110/120][]
- TEAC [WX-7000 Series][]
- TEAC [es8][]

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
[teac]: https://www.teac.co.jp/int/
[travis image]: http://img.shields.io/travis/gotmc/taffmat/master.svg
[travis link]: https://travis-ci.org/gotmc/taffmat
[WX-7000 Series]: http://teac-ipd.com/wx-7000/
