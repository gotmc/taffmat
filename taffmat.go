// Copyright (c) 2020–2024 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gotmc/convert"
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
	Filename    string
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

// ReadHeader reads the TAFFmat header file. The filename can, but is not
// required to, include the `.hdr` extension.
func ReadHeader(filename string) (*Header, error) {
	ext := filepath.Ext(filename)
	if strings.ToUpper(ext) == ".HDR" {
		filename = strings.TrimSuffix(filename, ext)
	} else if ext != "" {
		return nil, fmt.Errorf("filename extension must be blank or HDR instead of %s", ext)
	}
	data, err := os.ReadFile(filename + ".HDR")
	if err != nil {
		return nil, err
	}
	return parseHeader(data, filename)
}

func parseHeader(data []byte, filename string) (*Header, error) {
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

	hdr.Filename = filename
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
		return nil, fmt.Errorf("invalid device type: %s", hdrMap["device"])
	}
	hdr.Device = dt

	// Configure channels
	offsets, err := convert.StringToFloats(hdrMap["y_offset"], ",")
	if err != nil {
		return nil, err
	} else if len(offsets) != hdr.NumSeries {
		return nil, fmt.Errorf("found %d y-offsets and %d series", len(offsets), hdr.NumSeries)
	}
	slopes, err := convert.StringToFloats(hdrMap["slope"], ",")
	if err != nil {
		return nil, err
	} else if len(slopes) != hdr.NumSeries {
		return nil, fmt.Errorf("found %d slopes and %d series", len(slopes), hdr.NumSeries)
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
