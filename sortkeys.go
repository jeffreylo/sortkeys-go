package main

import (
	"bufio"
	"flag"
	"go/token"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type config struct {
	// Filename defines the name of the file to be parsed.
	Filename string

	// OutputFilename defines the location of the output file.
	OutputFilename string

	// WriteToFile, if true, writes the formatted output to OutputFilename.
	WriteToFile bool

	preferredFields map[string]int
	file            *dst.File
}

// Parse validates the configuration and creates the parsed *dst.File.
func (c *config) Parse(ff string) error {
	file, err := os.OpenFile(c.Filename, os.O_RDWR, 0755)
	defer checkClose(file)
	if err != nil {
		return err
	}

	preferred := make(map[string]int)
	for i, fieldName := range strings.Split(ff, ",") {
		if fieldName != "" {
			preferred[fieldName] = i + 1
		}
	}
	c.preferredFields = preferred

	if c.OutputFilename == "" {
		c.OutputFilename = c.Filename
	}

	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, "", bufio.NewReader(file), 0)
	if err != nil {
		return err
	}
	c.file = f

	return nil
}

// Rewrite walks the AST and rewrite structs and interfaces with sorted
// field order.
func (c *config) Rewrite() error {
	dst.Inspect(c.file, func(n dst.Node) bool {
		switch x := n.(type) {
		case *dst.InterfaceType:
			sort.Sort(byFieldName{
				Fields:          x.Methods.List,
				PreferredFields: c.preferredFields,
			})
			setFieldDecorators(x.Methods.List)
		case *dst.StructType:
			sort.Sort(byFieldName{
				Fields:          x.Fields.List,
				PreferredFields: c.preferredFields,
			})
			setFieldDecorators(x.Fields.List)
		}
		return true
	})
	return nil
}

// Write outputs the file to the configured writer.
func (c *config) Write() error {
	var writer io.Writer
	if c.WriteToFile {
		file, err := os.OpenFile(c.OutputFilename, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	w := bufio.NewWriter(writer)
	defer w.Flush()
	return decorator.Fprint(w, c.file)
}

func main() {
	flagFilename := flag.String("file", "", "Filename to be parsed")
	flagOutputFilename := flag.String("o", "", "Output filename")
	flagPreferredFields := flag.String("p", "", "Comma-separated ordered list of field names to prioritize")
	flagWrite := flag.Bool("w", false, "Write to -file")
	flag.Parse()

	cfg := &config{
		Filename:       *flagFilename,
		OutputFilename: *flagOutputFilename,
		WriteToFile:    *flagWrite,
	}

	must(cfg.Parse(*flagPreferredFields))
	must(cfg.Rewrite())
	must(cfg.Write())
}

// byFieldName implements sort.Interface for []*dst.Field based on
// the field's name.
type byFieldName struct {
	Fields          []*dst.Field
	PreferredFields map[string]int
}

func (a byFieldName) Len() int      { return len(a.Fields) }
func (a byFieldName) Swap(i, j int) { a.Fields[i], a.Fields[j] = a.Fields[j], a.Fields[i] }
func (a byFieldName) Less(i, j int) bool {
	x := a.Fields[i].Names[0].Name
	y := a.Fields[j].Names[0].Name
	x1, ok1 := a.PreferredFields[x]
	y1, ok2 := a.PreferredFields[y]

	if ok1 && ok2 {
		return x1 < y1
	} else if ok1 && !ok2 {
		return true
	} else if !ok1 && ok2 {
		return false
	}
	return x < y
}

// setFieldDecorators manages the commentary around fields.
func setFieldDecorators(fields []*dst.Field) {
	for i, field := range fields {
		// If the field has a leading comment, add newlines as
		// appropriate excl. the first field.
		if len(field.Decs.Start.All()) > 0 {
			if i != 0 {
				field.Decs.Before = dst.EmptyLine
			} else {
				field.Decs.Before = dst.NewLine
			}
			field.Decs.After = dst.EmptyLine
		} else {
			field.Decs.Before = dst.NewLine
			field.Decs.After = dst.NewLine
		}
	}
}

// must exits the process if err is not nil.
func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// checkClose closes the io.Closer.
func checkClose(closer io.Closer) {
	must(closer.Close())
}
