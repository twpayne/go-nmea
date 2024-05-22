package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/flarm"
	"github.com/twpayne/go-nmea/garmin"
	"github.com/twpayne/go-nmea/lxnavigation"
	"github.com/twpayne/go-nmea/samsung"
	"github.com/twpayne/go-nmea/standard"
	"github.com/twpayne/go-nmea/ublox"
)

var sentenceRx = regexp.MustCompile(`\$[A-Z]+,[^*]+\*(?:[0-9A-Fa-f]{2})?`)

func processReader(encoder *json.Encoder, parser *nmea.Parser, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		match := sentenceRx.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		var value any
		switch sentence, err := parser.ParseString(match[0]); {
		case err == nil:
			value = map[string]any{
				sentence.GetAddress().String(): sentence,
			}
		default:
			value = map[string]any{
				"err":      err.Error(),
				"sentence": match[0],
			}
		}
		if err := encoder.Encode(value); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func processFile(encoder *json.Encoder, parser *nmea.Parser, name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return processReader(encoder, parser, file)
}

func run() error {
	outputFilename := flag.String("o", "", "output filename")
	flag.Parse()

	parser := nmea.NewParser(
		nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
		nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
		nmea.WithSentenceParserFunc(flarm.SentenceParserFunc),
		nmea.WithSentenceParserFunc(garmin.SentenceParserFunc),
		nmea.WithSentenceParserFunc(lxnavigation.SentenceParserFunc),
		nmea.WithSentenceParserFunc(samsung.SentenceParserFunc),
		nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		nmea.WithSentenceParserFunc(ublox.SentenceParserFunc),
	)

	var output *os.File
	if *outputFilename == "" || *outputFilename == "-" {
		output = os.Stdout
	} else {
		outputFile, err := os.Create(*outputFilename)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		output = outputFile
	}

	encoder := json.NewEncoder(output)

	if flag.NArg() == 0 {
		if err := processReader(encoder, parser, os.Stdin); err != nil {
			return err
		}
	} else {
		for _, arg := range flag.Args() {
			if err := processFile(encoder, parser, arg); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
