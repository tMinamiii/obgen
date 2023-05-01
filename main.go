package main

import (
	"flag"
	"log"

	"github.com/tMinamiii/datagen/datagen"
)

func main() {
	source, destination, format := parseFlag()

	gen := datagen.NewDataGen(format)

	if err := gen.Parse(source, destination); err != nil {
		log.Fatal(err)
	}
}

func parseFlag() (source, destination string, format datagen.DataFormat) {
	f := flag.String("format", "", "data format (yaml or json)")
	s := flag.String("source", "", "source go file")
	d := flag.String("destination", "", "output destination")
	flag.Parse()

	if f == nil || s == nil || d == nil {
		log.Fatal("format or source or destination is nil")
	} else if *f == "" || *s == "" || *d == "" {
		log.Fatal("format or source or destination is empty")
	} else if *f != datagen.JSON.String() || *f != datagen.YAML.String() {
		log.Fatalf("invalid format %s\n", *f)
	}

	return *s, *d, datagen.DataFormat(*f)
}
