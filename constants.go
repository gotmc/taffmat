// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

// DeviceType are the available Teac device types.
type DeviceType int

// Available device types.
const (
	LX10 DeviceType = iota
	LX20
	LX110
	LX120
)

var deviceMap = map[string]DeviceType{
	"LX-10":  LX10,
	"LX-20":  LX20,
	"LX-110": LX110,
	"LX-120": LX120,
}

var deviceDesc = map[DeviceType]string{
	LX10:  "LX-10",
	LX20:  "LX-20",
	LX110: "LX-110",
	LX120: "LX-120",
}

// String implements the Stringer interface for the DeviceType.
func (dt DeviceType) String() string {
	return deviceDesc[dt]
}

// FileType are the available Teac file types.
type FileType int

// Available device types.
const (
	IntegerFile FileType = iota
	LongFile
)

var fileMap = map[string]FileType{
	"INTEGER": IntegerFile,
	"LONG":    LongFile,
}

var fileDesc = map[FileType]string{
	IntegerFile: "Integer (16-bit A/D, 2-byte integers)",
	LongFile:    "Long (24-bit A/D, 4-byte integers)",
}

// String implements the Stringer interface for the DeviceType.
func (ft FileType) String() string {
	return fileDesc[ft]
}

// BitResolution returns the bit resolution of the Analog-to-Digital Converter
// (ADC).
func (ft FileType) BitResolution() int {
	if ft == IntegerFile {
		return 16
	} else if ft == LongFile {
		return 24
	}
	return 0
}

// NumBytes returns the number of bytes in each integer.
func (ft FileType) NumBytes() int {
	if ft == IntegerFile {
		return 2
	} else if ft == LongFile {
		return 4
	}
	return 0
}
