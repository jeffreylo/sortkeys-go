package main

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/deliveroo/assert-go"
)

var update = flag.Bool("update", false, "update golden files")

const golden = "testdata/golden/example.go"

func TestSortKeys(t *testing.T) {
	var file *os.File

	if *update {
		f, _ := os.OpenFile(golden, os.O_RDWR, 0755)
		defer checkClose(f)
		file = f
	} else {
		f, _ := ioutil.TempFile("", "prefix")
		defer func() {
			os.Remove(file.Name())
		}()
		file = f
	}

	cfg := &config{
		Filename:       "testdata/fixtures/example.go",
		OutputFilename: file.Name(),
		WriteToFile:    true,
	}
	must(cfg.Parse("ID"))
	must(cfg.Rewrite())
	must(cfg.Write())

	got, _ := ioutil.ReadFile(file.Name())
	want, _ := ioutil.ReadFile(golden)
	assert.Equal(t, string(got), string(want))
}
