package main

import (
	"bytes"
	"log"
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
		offset0  int64
		limit0   int64
		len      int64
		offset   int64
		limit    int64
	}{
		{"testdata/input.txt", "testdata/out_offset0_limit0.txt", "t0.txt", 0, 0, 6742, 0, 0},
		{"testdata/input.txt", "testdata/out_offset0_limit10.txt", "t1.txt", 0, 0, 11, 0, 10},
		{"testdata/input.txt", "testdata/out_offset0_limit1000.txt", "t2.txt", 0, 0, 1026, 0, 1000},
		{"testdata/input.txt", "testdata/out_offset0_limit10000.txt", "t3.txt", 0, 0, 6742, 0, 10000},
		{"testdata/input.txt", "testdata/out_offset6000_limit1000.txt", "t5.txt", 0, 0, 629, 6000, 1000},
	}
	for i := range testlen {
		tc := testlen[i]
		t.Run("test len", func(t *testing.T) {
			err := Copy(tc.idelFile, strconv.Itoa(i)+".txt", tc.offset, tc.limit)
			if err != nil {
				require.Error(t, err, "offset > limit")
			} else {
				if err := Copy(tc.pathfrom, tc.pathto, tc.offset0, tc.limit); err != nil {
					log.Fatal(err)
				}
				file, err := os.Open(tc.pathto)
				require.NoError(t, err)
				siz, err := file.Stat()
				require.NoError(t, err)
				require.Nil(t, err, "ok")
				ss := siz.Size
				if tc.limit != 0 {
					require.Equal(t, tc.limit, ss())
				} else {
					require.Equal(t, tc.len, ss())
				}
			}
		})
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
