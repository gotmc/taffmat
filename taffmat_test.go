// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import "testing"

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
Y_OFFSET 0.000000e+000,0.000000e+000
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
SERIES CH1_LX10_DC100K,CH2_LX10_DC100K,CH3_LX10_DC100K,CH4_LX10_DC100K,CH5_LX110_DC100K,CH6_LX10_DC100K,CH7_LX10_DC100K,CH8_LX10_DC100K DATE 06-23-2001
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
		hdr         string
		dataset     string
		fileVersion int
		numSeries   int
		numSamples  int
		storageMode StorageType
		xOffset     float64
		memo        string
		device      DeviceType
		fileType    FileType
	}{
		{
			hdr:         lx10Header1,
			dataset:     "UTEST001",
			fileVersion: 1,
			numSeries:   2,
			numSamples:  1249792,
			storageMode: Interlaced,
			xOffset:     0.0,
			memo:        "Sample recordings for unit_testing taffmat.py",
			device:      LX10,
			fileType:    IntegerFile,
		},
		{
			hdr:         lx10Header2,
			dataset:     "TEST0001",
			fileVersion: 1,
			numSeries:   8,
			numSamples:  59200,
			storageMode: Interlaced,
			xOffset:     -5,
			memo:        "",
			device:      LX10,
			fileType:    IntegerFile,
		},
		{
			hdr:         lx110Header1,
			dataset:     "TEST0001",
			fileVersion: 1,
			numSeries:   8,
			numSamples:  59200,
			storageMode: Interlaced,
			xOffset:     -5.0,
			memo:        "<LX-110>",
			device:      LX110,
			fileType:    IntegerFile,
		},
	}
	for _, tc := range testCases {
		hdr, err := parseHeader([]byte(tc.hdr))
		if err != nil {
			t.Errorf("error parsing file: %s", err)
		}
		if hdr.Dataset != tc.dataset {
			t.Errorf("Dataset = %s, expected %s", hdr.Dataset, tc.dataset)
		}
		if hdr.FileVersion != tc.fileVersion {
			t.Errorf("FileVersion = %d, expected %d", hdr.FileVersion, tc.fileVersion)
		}
		if hdr.NumSeries != tc.numSeries {
			t.Errorf("NumSeries = %d, expected %d", hdr.NumSeries, tc.numSeries)
		}
		if hdr.NumSamples != tc.numSamples {
			t.Errorf("NumSamples = %d, expected %d", hdr.NumSamples, tc.numSamples)
		}
		if hdr.StorageMode != tc.storageMode {
			t.Errorf("StorageMode = %s, expected %s", hdr.StorageMode, tc.storageMode)
		}
		if hdr.XOffset != tc.xOffset {
			t.Errorf("XOffset = %f, expected %f", hdr.XOffset, tc.xOffset)
		}
		if hdr.Memo != tc.memo {
			t.Errorf("Memo = %s, expected %s", hdr.Memo, tc.memo)
		}
		if hdr.Device != tc.device {
			t.Errorf("Device = %s, expected %s", hdr.Device, tc.device)
		}
		if hdr.FileType != tc.fileType {
			t.Errorf("FileType = %s, expected %s", hdr.FileType, tc.fileType)
		}
	}
}
