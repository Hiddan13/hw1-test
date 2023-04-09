package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset > limit {
		return fmt.Errorf("offset > limit")
	}
	buf := make([]byte, limit)
	if limit == 0 {
		b, err := os.ReadFile(fromPath)
		if err != nil {
			fmt.Println(err)
		}
		buf = b
	}
	file, err := os.Open(fromPath)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		return err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fSize := fileInfo.Size

	if err != nil {
		return err
	}
	if offset < fSize() {
		file.Seek(offset, io.SeekStart)
		file.Read(buf)
		f, err := os.Create(toPath)
		if err != nil {
			return err
		}
		w, err := f.Write(buf)
		if err != nil && w >= 0 {
			return err
		}
	}
	return nil
}
