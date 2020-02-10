// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package taffmat reads and writes TAFFmat file.  TAFFmat is TEAC's proprietary
file format used to store data from their LX series and other data recorders.

According to the TEAC "LX Series Recording Unit Instruction Manual":

>  TAFFmat (an acronym for TEAC Data Acquisition File Format) is a
>  file format composed of the following:
>
>  * a data file containing A/D (analog to digital) converted data. The
>    file is binary format with the extension dat.
>  * a header file containing information such as recording
>    conditions. The file is in text format with the extension hdr.

TAFFmat is a trademark of TEAC Corporation.

The following data recorders store their data in the TAFFmat file format:

- TEAC [LX-10/20][]
- TEAC [LX-110/120][]
- TEAC [WX-7000 Series][]
- TEAC [es8][]
*/
package taffmat
