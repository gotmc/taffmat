// Copyright (c) 2020–2023 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package taffmat reads and writes TAFFmat (TEAC Data Acquisition File Format)
file.  TAFFmat is TEAC's proprietary file format used to store data from their
LX series and other data recorders.

TAFFmat files are composed of two files. One is a binary data file with the
extension `dat` containing A/D (analog to digital) converted data. The other a
text-based header file with the extension `hdr` containing information such as
recording conditions.
*/
package taffmat
