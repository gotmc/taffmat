// Copyright (c) 2020 The taffmat developers. All rights reserved.
// Project site: https://github.com/gotmc/taffmat
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package taffmat

import "testing"

const givenHdr = `DATASET UTEST001
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

const givenHdr2 = `DATASET TEST0001
VERSION 1
SERIES CH1_LX10_DC100K,CH2_LX10_DC100K,CH3_LX10_DC100K,CH4_LX10_DC100K, CH5_LX10_DC100K,CH6_LX10_DC100K,CH7_LX10_DC100K,CH8_LX10_DC100K
DATE 06-23-2001
TIME 16:32:55.00
RATE 1500
VERT_UNITS V,V,V,V,V,V,V,V HORZ_UNITS Sec
COMMENT <LX-10> NUM_SERIES 8 STORAGE_MODE INTERLACED FILE_TYPE INTEGER
SLOPE 0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000,0.00008000 X_OFFSET -5
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
STOP_CONDITION COMMAND,POST START_PRE_COUNT 7500
STOP_POST_COUNT 15000
MARK 100,200,300
ID_END
VOICE_MEMO 8BITS,327680
LX10_VERSION PAL1_VER,PAL2_VER,V0.03,02200000
`

func TestParseHeader(t *testing.T) {
	hdr, err := parseHeader([]byte(givenHdr))
	if err != nil {
		t.Errorf("error parsing file: %s", err)
	}
	if hdr.Dataset != "UTEST001" {
		t.Errorf("Dataset = %s, expected UTEST001", hdr.Dataset)
	}
	if hdr.Version != 1 {
		t.Errorf("Version = %d, expected 1", hdr.Version)
	}
	if hdr.NumSeries != 2 {
		t.Errorf("NumSeries = %d, expected 2", hdr.NumSeries)
	}
	if hdr.NumSamples != 1249792 {
		t.Errorf("NumSamples = %d, expected 1249792", hdr.NumSamples)
	}
	if hdr.StorageMode != "INTERLACED" {
		t.Errorf("StorageMode = %s, expected INTERLACED", hdr.StorageMode)
	}
	if hdr.XOffset != 0.0 {
		t.Errorf("XOffset = %f, expected 0.0", hdr.XOffset)
	}
	expectedMemo := "Sample recordings for unit_testing taffmat.py"
	if hdr.Memo != expectedMemo {
		t.Errorf("Memo = %s, expected %s", hdr.Memo, expectedMemo)
	}
	expectedDevice := LX10
	if hdr.Device != expectedDevice {
		t.Errorf("Device = %s, expected %s", hdr.Device, expectedDevice)
	}
	expectedFileType := IntegerFile
	if hdr.FileType != expectedFileType {
		t.Errorf("FileType = %s, expected %s", hdr.FileType, expectedFileType)
	}
}