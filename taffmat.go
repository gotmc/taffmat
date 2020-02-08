// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import (
	"fmt"
	"io/ioutil"
	"log"
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
			log.Printf("hdrMap[%s] = %s", key, val)
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

	// Determine XOffset
	xOffset, err := strconv.ParseFloat(hdrMap["x_offset"], 32)
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

	// Determine Memo
	hdr.Memo = hdrMap["memo"]

	return &hdr, nil
}
