package main

import (
	"bytes"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	testlen := []struct {
		idelFile string
		pathfrom string
		pathto   string
		len      int64
		offset   int64
		limit    int64
	}{
		{"testdata/input.txt", "testdata/out_offset0_limit0.txt", "t0.txt", 6742, 0, 0},
		{"testdata/input.txt", "testdata/out_offset0_limit10.txt", "t1.txt", 11, 0, 10},
		{"testdata/input.txt", "testdata/out_offset0_limit1000.txt", "t2.txt", 1026, 0, 1000},
		{"testdata/input.txt", "testdata/out_offset0_limit10000.txt", "t3.txt", 6742, 0, 10000},
		{"testdata/input.txt", "testdata/out_offset6000_limit1000.txt", "t5.txt", 629, 6000, 1000},
	}
	for i := range testlen {
		tc := testlen[i]
		t.Run("offset and limit is OK", func(t *testing.T) {
			err := Copy(tc.idelFile, strconv.Itoa(i)+".txt", tc.offset, tc.limit)
			if err != nil {
				require.Error(t, err, "offset > limit")
			} else {
				fileTest, err := os.Open(strconv.Itoa(i) + ".txt")
				require.NoError(t, err)
				bufTest := make([]byte, tc.limit)
				fileTest.Read(bufTest)
				fileData, err := os.Open(tc.pathfrom)
				require.NoError(t, err)
				bufData := make([]byte, tc.limit)
				fileData.Read(bufData)
				require.True(t, bytes.Equal(bufData, bufTest), "OK")
			}
		})
	}
}
