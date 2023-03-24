package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	// Place your code here.
}
func CCopy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	if offset > limit {
		return fmt.Errorf("offset > limit")
	}
	buf := make([]byte, limit)
	file, err := os.Open(fromPath)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if limit == 0 {
		b, err := os.ReadFile(fromPath)
		if err != nil {
			fmt.Println(err)
		}
		buf = b
	}

	//io.Read()

	if err != nil {
		return err
	}
	siz, err := file.Stat()
	if err != nil {
		return err
	}
	ss := siz.Size
	//fmt.Println(ss()) // получим длину файла в байтах
	if err != nil {
		return err
	}
	if offset < ss() {
		file.Seek(offset, io.SeekStart) // на тчениие указываем что читаем с отступом
		file.Read(buf)
		f, err := os.Create(toPath)
		if err != nil {
			return err
		}
		w, _ := f.Write(buf) // сколько в итоге хаписали
		fmt.Println(w)       // сколько в итоге хаписали
		//fmt.Println("ok")    // сколько в итоге хаписали
	}
	return nil

}
