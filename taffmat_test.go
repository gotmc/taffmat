// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import (
	"math"
	"testing"
	"time"
)

const lx10Header1 = `DATASET UTEST001
VERSION 1
SERIES CH1_LX-10_DC100K,CH2_LX-10_DC100K
DATE 02-09-2013
TIME 13:35:37.00
RATE 96000
VERT_UNITS V,V
HORZ_UNITS Sec
COMMENT Sample recordings for unit_testing taffmat.py
NUM_SERIES 2
STORAGE_MODE INTERLACED
FILE_TYPE INTEGER
SLOPE 8.000000e-005,2.000000e-004
X_OFFSET 0.0
Y_OFFSET 0.100000e+000,0.000000e+000
NUM_SAMPS 1249792
DATA
DEVICE LX-10
SLOT1_AMP AD_AMP,8,V01.21  ,V02.34
SLOT2_AMP DA_AMP,8,V02.19  ,
CH1_1 LX-10_DC100K,RANGE=2V,FILTER=ON
CH2_2 LX-10_DC100K,RANGE=5V,FILTER=ON
ID_NO 1
TIME 20130209133537,20130209133550
REC_MODE PCCARD
START_TRIGGER COMMAND
STOP_CONDITION COMMAND
ID_END
LX10_VERSION V01.09  ,V01.14  ,V5.02g  ,00022E202239
MEMO_LENGTH 45,0,0,0,0,0,0,0
MEMO Sample recordings for unit_testing taffmat.py

 `

const lx10Header2 = `DATASET TEST0001
VERSION 1
SERIES CH1_LX10_DC100K,CH2_LX10_DC100K,CH3_LX10_DC100K,CH4_LX10_DC100K,CH5_LX10_DC100K,CH6_LX10_DC100K,CH7_LX10_DC100K,CH8_LX10_DC100K
DATE 06-23-2001
TIME 16:32:55.00
RATE 1500
VERT_UNITS V,V,V,V,V,V,V,V
HORZ_UNITS Sec
COMMENT <LX-10>
NUM_SERIES 8
STORAGE_MODE INTERLACED
FILE_TYPE INTEGER
SLOPE 0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000
X_OFFSET -5
Y_OFFSET 0,0,0,0,0,0,0,0
NUM_SAMPS 59200
DATA
DEVICE LX-10
SLOT1_AMP AD_AMP,8,00000006,
SLOT2_AMP DA_AMP,8,00000006,
CH1_1 LX10_DC100K,RANGE=2V,FILTER=ON
CH2_2 LX10_DC100K,RANGE=2V,FILTER=ON
CH3_3 LX10_DC100K,RANGE=2V,FILTER=ON
CH4_4 LX10_DC100K,RANGE=2V,FILTER=ON
CH5_5 LX10_DC100K,RANGE=2V,FILTER=ON
CH6_6 LX10_DC100K,RANGE=2V,FILTER=ON
CH7_7 LX10_DC100K,RANGE=2V,FILTER=ON
CH8_8 LX10_DC100K,RANGE=2V,FILTER=ON
ID_NO 1
TIME 20010623163255,20010623163335
REC_MODE MO
START_TRIGGER COMMAND,PRE
STOP_CONDITION COMMAND,POST
START_PRE_COUNT 7500
STOP_POST_COUNT 15000
MARK 100,200,300
ID_END
VOICE_MEMO 8BITS,327680
LX10_VERSION PAL1_VER,PAL2_VER,V0.03,02200000
`

const lx110Header1 = `DATASET TEST0001
VERSION 1
SERIES CH1_LX110_DC100K,CH2_LX110_DC100K,CH3_LX110_DC100K,CH4_LX110_DC100K,CH5_LX110_DC100K,CH6_LX110_DC100K,CH7_LX110_DC100K,CH8_LX110_DC100K
DATE 06-23-2001
TIME 16:32:55.00
RATE 1500
VERT_UNITS V,V,V,V,V,V,V,V
HORZ_UNITS Sec
COMMENT <LX-110>
NUM_SERIES 8
STORAGE_MODE INTERLACED
FILE_TYPE INTEGER
SLOPE 8.000000e-005, 8.000000e-005, 8.000000e-005, 8.000000e-005, 8.000000e-005, 8.000000e-005, 8.000000e-005, 8.000000e-005
X_OFFSET -5.0
Y_OFFSET 0.000000e+000, 0.000000e+000, 0.000000e+000, 0.000000e+000, 0.000000e+000, 0.000000e+000, 0.000000e+000, 0.000000e+000
NUM_SAMPS 59200
DATA
DEVICE LX-110
SLOT1_AMP AD_AMP,8,V1.00 ,V1.00 SLOT2_AMP DA_AMP,8,V1.00 ,V1.00
CH1_1 LX110_DC100K,RANGE=2V,FILTER=ON
CH2_2 LX110_DC100K,RANGE=2V,FILTER=ON
CH3_3 LX110_DC100K,RANGE=2V,FILTER=ON
CH4_4 LX110_DC100K,RANGE=2V,FILTER=ON
CH5_5 LX110_DC100K,RANGE=2V,FILTER=ON
CH6_6 LX110_DC100K,RANGE=2V,FILTER=ON
CH7_7 LX110_DC100K,RANGE=2V,FILTER=ON
CH8_8 LX110_DC100K,RANGE=2V,FILTER=ON
ID_NO 1
TIME 20010623163255,20010623163335
REC_MODE MO
START_TRIGGER COMMAND,PRE
STOP_CONDITION COMMAND,POST
START_PRE_COUNT 7500
STOP_POST_COUNT 15000
MARK 100,200,300
ID_END
VOICE_MEMO 8BITS,327680
LX110_VERSION V1.00 ,V1.00 ,V1.00 ,00022E202000
MEMO_LENGTH 8,0,0,0,0,0,0,0
MEMO <LX-110>`

func TestParseHeader(t *testing.T) {
	testCases := []struct {
		given string
		hdr   Header
	}{
		{
			given: lx10Header1,
			hdr: Header{
				Dataset:     "UTEST001",
				FileVersion: 1,
				StartTime:   time.Date(2013, 2, 9, 13, 35, 37, 0, time.UTC),
				StopTime:    time.Date(2013, 2, 9, 13, 35, 50, 0, time.UTC),
				NumSeries:   2,
				NumSamples:  1249792,
				StorageMode: Interlaced,
				XOffset:     0.0,
				Memo:        "Sample recordings for unit_testing taffmat.py",
				Device:      LX10,
				FileType:    IntegerFile,
				Channels: []Channel{
					{
						Number:  1,
						Slope:   0.00008,
						YOffset: 0.1,
					},
					{
						Number:  2,
						Slope:   0.0002,
						YOffset: 0.0,
					},
				},
			},
		},
		{
			given: lx10Header2,
			hdr: Header{
				Dataset:     "TEST0001",
				FileVersion: 1,
				StartTime:   time.Date(2001, 6, 23, 16, 32, 55, 0, time.UTC),
				StopTime:    time.Date(2001, 6, 23, 16, 33, 35, 0, time.UTC),
				NumSeries:   8,
				NumSamples:  59200,
				StorageMode: Interlaced,
				XOffset:     -5,
				Memo:        "",
				Device:      LX10,
				FileType:    IntegerFile,
				Channels: []Channel{
					{
						Number: 1,
						Slope:  0.00008,
					},
					{
						Number: 2,
						Slope:  0.00008,
					},
				},
			},
		},
		{
			given: lx110Header1,
			hdr: Header{
				Dataset:     "TEST0001",
				FileVersion: 1,
				StartTime:   time.Date(2001, 6, 23, 16, 32, 55, 0, time.UTC),
				StopTime:    time.Date(2001, 6, 23, 16, 33, 35, 0, time.UTC),
				NumSeries:   8,
				NumSamples:  59200,
				StorageMode: Interlaced,
				XOffset:     -5.0,
				Memo:        "<LX-110>",
				Device:      LX110,
				FileType:    IntegerFile,
				Channels: []Channel{
					{
						Number: 1,
						Slope:  0.00008,
					},
					{
						Number: 2,
						Slope:  0.00008,
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		hdr, err := parseHeader([]byte(tc.given), "foo")
		if err != nil {
			t.Errorf("error parsing file: %s", err)
		}
		if hdr.Dataset != tc.hdr.Dataset {
			t.Errorf("Dataset = %s, expected %s", hdr.Dataset, tc.hdr.Dataset)
		}
		if hdr.FileVersion != tc.hdr.FileVersion {
			t.Errorf("FileVersion = %d, expected %d", hdr.FileVersion, tc.hdr.FileVersion)
		}
		if hdr.NumSeries != tc.hdr.NumSeries {
			t.Errorf("NumSeries = %d, expected %d", hdr.NumSeries, tc.hdr.NumSeries)
		}
		if hdr.NumSamples != tc.hdr.NumSamples {
			t.Errorf("NumSamples = %d, expected %d", hdr.NumSamples, tc.hdr.NumSamples)
		}
		if hdr.StorageMode != tc.hdr.StorageMode {
			t.Errorf("StorageMode = %s, expected %s", hdr.StorageMode, tc.hdr.StorageMode)
		}
		if hdr.XOffset != tc.hdr.XOffset {
			t.Errorf("XOffset = %f, expected %f", hdr.XOffset, tc.hdr.XOffset)
		}
		if hdr.Memo != tc.hdr.Memo {
			t.Errorf("Memo = %s, expected %s", hdr.Memo, tc.hdr.Memo)
		}
		if hdr.Device != tc.hdr.Device {
			t.Errorf("Device = %s, expected %s", hdr.Device, tc.hdr.Device)
		}
		if hdr.FileType != tc.hdr.FileType {
			t.Errorf("FileType = %s, expected %s", hdr.FileType, tc.hdr.FileType)
		}
		if hdr.StartTime != tc.hdr.StartTime {
			t.Errorf("StartTime = %s, expected %s", hdr.StartTime, tc.hdr.StartTime)
		}
		if hdr.StopTime != tc.hdr.StopTime {
			t.Errorf("StopTime = %s, expected %s", hdr.StopTime, tc.hdr.StopTime)
		}
		if len(hdr.Channels) != hdr.NumSeries {
			t.Errorf("Channel length (%d) doesn't match number of series (%d)", len(hdr.Channels), hdr.NumSeries)
		}
		if !almostEqual(hdr.Channels[0].Slope, tc.hdr.Channels[0].Slope) {
			t.Errorf("slope = %f, expected %f", hdr.Channels[0].Slope, tc.hdr.Channels[0].Slope)
		}
		if !almostEqual(hdr.Channels[1].Slope, tc.hdr.Channels[1].Slope) {
			t.Errorf("slope = %f, expected %f", hdr.Channels[1].Slope, tc.hdr.Channels[1].Slope)
		}
		if !almostEqual(hdr.Channels[0].YOffset, tc.hdr.Channels[0].YOffset) {
			t.Errorf("ch 1 y-offset = %f, expected %f", hdr.Channels[0].YOffset, tc.hdr.Channels[0].YOffset)
		}
		if !almostEqual(hdr.Channels[1].YOffset, tc.hdr.Channels[1].YOffset) {
			t.Errorf("ch 2 y-offset = %f, expected %f", hdr.Channels[1].YOffset, tc.hdr.Channels[1].YOffset)
		}
	}
}

const tolerance = 0.0000000001

func almostEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < tolerance
}
