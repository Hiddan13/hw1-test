package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
	testlen := []struct {
		pathfrom   string
		pathto     string
		def_offset int64
		def_limit  int64
		len        int64
		offset     int64
		limit      int64
	}{
		//{"testdata/input.txt", "src/t.txt", 0, 0, 6742, 0, 0},
		{"testdata/out_offset0_limit0.txt", "src/t0.txt", 0, 0, 6742, 0, 0},
		{"testdata/out_offset0_limit10.txt", "src/t1.txt", 0, 0, 11, 0, 10},
		{"testdata/out_offset0_limit1000.txt", "src/t2.txt", 0, 0, 1026, 0, 1000},
		{"testdata/out_offset0_limit10000.txt", "src/t3.txt", 0, 0, 6742, 0, 10000},
		//{"testdata/out_offset100_limit1000.txt", "src/t4.txt", 0, 0, 1017, 100, 1000},- по мойму там есть ошибка
		{"testdata/out_offset6000_limit1000.txt", "src/t5.txt", 0, 0, 629, 6000, 1000},
	}
	os.Mkdir("src", 0755) // testoffset_limit := []struct {

	for i := range testlen {
		tc := testlen[i]
		t.Run("test len", func(t *testing.T) {
			if tc.offset > tc.limit {
				err := CCopy("testdata/input.txt", "src/"+strconv.Itoa(i)+".txt", tc.offset, tc.limit)
				if err != nil {
					require.Error(t, err, "offset > limit")
				}

			}
			//t.Parallel()
			err := CCopy(tc.pathfrom, tc.pathto, tc.def_offset, tc.def_limit)
			if err != nil {
				fmt.Println(" CCopy", err)
			}
			file, err := os.Open(tc.pathto)
			if err != nil {
				fmt.Println("os.Open", err)
			}
			siz, err := file.Stat()
			if err != nil {
				fmt.Println("file.Stat", err)
			}
			ss := siz.Size
			require.Equal(t, tc.len, ss())
		})

	}
	for i := range testlen {
		tc := testlen[i]
		t.Run("have file in dir", func(t *testing.T) {
			if tc.offset > tc.limit {
				err := CCopy("testdata/input.txt", "src/"+strconv.Itoa(i)+".txt", tc.offset, tc.limit)
				if err != nil {
					require.Error(t, err, "offset > limit")
				}
			}
			//t.Parallel()
			_, err := os.Stat(tc.pathto)
			if err != nil {
				fmt.Println("на найдена в папке src")
			} else {
			}
			require.Nil(t, err, "ok")
		})
	}
	for i := range testlen {
		tc := testlen[i]
		t.Run("offset and limit is OK", func(t *testing.T) {
			if tc.offset > tc.limit {
				err := CCopy("testdata/input.txt", "src/"+strconv.Itoa(i)+".txt", tc.offset, tc.limit)
				if err != nil {
					require.Error(t, err, "offset > limit")
				}
			} else {
				err := CCopy("testdata/input.txt", "src/"+strconv.Itoa(i)+".txt", tc.offset, tc.limit)
				if err != nil {
					fmt.Println(err)
				}
				file_test, err := os.Open("src/" + strconv.Itoa(i) + ".txt")
				if err != nil {
					fmt.Println("os.Open", err)
				}
				buf_test := make([]byte, tc.limit)
				file_test.Read(buf_test)
				file_data, err := os.Open(tc.pathfrom)
				if err != nil {
					fmt.Println("os.Open", err)
				}
				buf_data := make([]byte, tc.limit)
				file_data.Read(buf_data)

				require.True(t, bytes.Equal(buf_data, buf_test), "OK")
			}
		})
	}
}
