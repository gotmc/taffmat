// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

// DeviceType are the available Teac device types.
type DeviceType string

// Available device types.
const (
	LX10  DeviceType = "LX-10"
	LX20             = "LX-20"
	LX110            = "LX-110"
	LX120            = "LX-120"
)

var deviceMap = map[string]DeviceType{
	"LX-10":  LX10,
	"LX-20":  LX20,
	"LX-110": LX110,
	"LX-120": LX120,
}

// String implements the Stringer interface for the DeviceType.
func (dt DeviceType) String() string {
	return string(dt)
}

// FileType are the available Teac file types.
type FileType string

// Available file types.
const (
	IntegerFile FileType = "INTEGER" // 16-bit ADC, 2-byte integers
	LongFile             = "LONG"    // 24-bit ADC, 4-byte integers
)

var fileMap = map[string]FileType{
	"INTEGER": IntegerFile,
	"LONG":    LongFile,
}

// String implements the Stringer interface for the DeviceType.
func (ft FileType) String() string {
	return string(ft)
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

// TriggerType are the available trigger types to start a recording.
type TriggerType int

// Available start recording trigger types.
const (
	CommandTrigger TriggerType = iota
	PanelTrigger
	LevelTrigger
	DateTrigger
	TimerTrigger
	ExternalTrigger
	TimeOutTrigger
)

var triggerMap = map[string]TriggerType{
	"COMMAND":  CommandTrigger,
	"PANEL":    PanelTrigger,
	"DATE":     DateTrigger,
	"TIMER":    TimerTrigger,
	"EXT":      ExternalTrigger,
	"TIME_OUT": TimeOutTrigger,
}

var triggerDesc = map[TriggerType]string{
	CommandTrigger:  "Interface command",
	PanelTrigger:    "FWD button of the front panel",
	DateTrigger:     "When Repeat Count is 1 in the interval action",
	TimerTrigger:    "When Repeat Count is 2 or more in the interval action",
	ExternalTrigger: "External Trigger",
	TimeOutTrigger:  "Time out",
}

// String implements the Stringer interface for the TriggerType.
func (tt TriggerType) String() string {
	return triggerDesc[tt]
}

// StorageType are the available storage types for the TAFFmat file.
type StorageType string

// Available TAFFmat file storage types.
const (
	Interlaced StorageType = "INTERLACED"
)

var storageMap = map[string]StorageType{
	"INTERLACED": Interlaced,
}

// String implements the Stringer interface for the StorageType.
func (st StorageType) String() string {
	return string(st)
}

// AmpType are the available TEAC amplifier types.
type AmpType string

// Available TEAC amplifier types.
const (
	DC100K AmpType = "DC100K"
)

var ampMap = map[string]AmpType{
	"DC100K": DC100K,
}

// String implements the Stringer interface for the AmpType.
func (at AmpType) String() string {
	return string(at)
}

// RangeType are the available TEAC amplifier types.
type RangeType string

// Available TEAC recording ranges.
const (
	Range2V RangeType = "2V"
	Range5V           = "5V"
)

var rangeMap = map[string]RangeType{
	"2V": Range2V,
	"5V": Range5V,
}

// String implements the Stringer interface for the RangeType.
func (rt RangeType) String() string {
	return string(rt)
}
