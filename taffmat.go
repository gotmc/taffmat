// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Amp models a Teac amplifier that can be in one of the two slots in the Teac
// LX recorder.
type Amp struct {
	Slot            int
	Name            string
	NumChannels     int
	VersionPLD      string
	VersionFirmware string
}

// Channel models a channel that recorded data.
type Channel struct {
	Name    string
	Number  int // 1 based channel number.
	Units   string
	Slope   float64
	YOffset float64
	Amp     AmpType
	Range   RangeType
	Filter  bool
}

// Header models the TAFFmat file header.
type Header struct {
	Dataset     string
	FileVersion int
	StartTime   time.Time
	StopTime    time.Time
	Rate        int
	NumSeries   int
	NumSamples  int
	StorageMode StorageType
	XOffset     float64
	Memo        string
	Device      DeviceType
	FileType    FileType
	Channels    []Channel
}

// ReadHeader reads the TAFFmat header file.
func ReadHeader(filename string) (*Header, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseHeader(data)
}

func parseHeader(data []byte) (*Header, error) {
	var hdr Header
	hdrMap := make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		keysValues := strings.SplitAfterN(line, " ", 2)
		if len(keysValues) == 2 {
			key := strings.TrimSpace(strings.ToLower(keysValues[0]))
			val := strings.TrimSpace(keysValues[1])
			hdrMap[key] = val
		} else if len(keysValues) == 1 {
			hdrMap[keysValues[0]] = ""
		}
	}

	hdr.Dataset = hdrMap["dataset"]

	// Determine FileVersion
	ver, err := strconv.Atoi(hdrMap["version"])
	if err != nil {
		return nil, err
	}
	if ver != 1 {
		return nil, fmt.Errorf("unknown version: %d", ver)
	}
	hdr.FileVersion = ver

	// Determine StorageMode
	storageMode := hdrMap["storage_mode"]
	if storageMode != "INTERLACED" {
		return nil, fmt.Errorf("unkown storage mode: %s", hdr.StorageMode)
	}
	hdr.StorageMode = Interlaced

	// Determine Rate
	rate, err := strconv.Atoi(hdrMap["rate"])
	if err != nil {
		return nil, err
	}
	hdr.Rate = rate

	// Determine NumSeries
	numSeries, err := strconv.Atoi(hdrMap["num_series"])
	if err != nil {
		return nil, err
	}
	hdr.NumSeries = numSeries

	// Determine NumSamples
	numSamples, err := strconv.Atoi(hdrMap["num_samps"])
	if err != nil {
		return nil, err
	}
	hdr.NumSamples = numSamples

	// Determine Filetype
	ft, ok := fileMap[hdrMap["file_type"]]
	if !ok {
		return nil, fmt.Errorf("invalid file type: %s", hdrMap["file_type"])
	}
	hdr.FileType = ft

	// Determine Start and Stop times.
	timeLayout := "20060102150405"
	timeStrings := strings.Split(hdrMap["time"], ",")
	startTime, err := time.Parse(timeLayout, timeStrings[0])
	if err != nil {
		return nil, err
	}
	hdr.StartTime = startTime
	stopTime, err := time.Parse(timeLayout, timeStrings[1])
	if err != nil {
		return nil, err
	}
	hdr.StopTime = stopTime

	// Determine XOffset
	xOffset, err := strconv.ParseFloat(hdrMap["x_offset"], 64)
	if err != nil {
		return nil, err
	}
	hdr.XOffset = xOffset

	// Determine Device
	dt, ok := deviceMap[hdrMap["device"]]
	if !ok {
		return nil, fmt.Errorf("Invalid device type: %s", hdrMap["device"])
	}
	hdr.Device = dt

	// Configure channels
	offsetStrings := strings.Split(hdrMap["y_offset"], ",")
	if len(offsetStrings) != hdr.NumSeries {
		return nil, fmt.Errorf("found %d y-offsets and %d series", len(offsetStrings), hdr.NumSeries)
	}
	offsets, err := parseStringFloat(offsetStrings)
	if err != nil {
		return nil, err
	}
	slopeStrings := strings.Split(hdrMap["slope"], ",")
	if len(slopeStrings) != hdr.NumSeries {
		return nil, fmt.Errorf("found %d slopes and %d series", len(slopeStrings), hdr.NumSeries)
	}
	slopes, err := parseStringFloat(slopeStrings)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= hdr.NumSeries; i++ {
		chKey := fmt.Sprintf("ch%d_%d", i, i)
		chString, ok := hdrMap[chKey]
		if !ok {
			return nil, fmt.Errorf("error finding CH%d_%d", i, i)
		}
		chInfo := strings.Split(chString, ",")
		rangeString := chInfo[1]
		rngStringVal := strings.Split(rangeString, "=")
		rngVal, ok := rangeMap[rngStringVal[1]]
		if !ok {
			return nil, fmt.Errorf("error finding range value in %s", rangeString)
		}
		ch := Channel{
			Name:    chInfo[0],
			Number:  i,
			Range:   rngVal,
			YOffset: offsets[i-1],
			Slope:   slopes[i-1],
		}
		hdr.Channels = append(hdr.Channels, ch)
	}

	// Determine Memo
	hdr.Memo = hdrMap["memo"]

	return &hdr, nil
}

func parseStringFloat(slice []string) ([]float64, error) {
	nums := make([]float64, len(slice))
	for i, s := range slice {
		num, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to float64", s)
		}
		nums[i] = num
	}
	return nums, nil
}
