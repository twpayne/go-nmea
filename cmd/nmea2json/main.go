package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/garmin"
	"github.com/twpayne/go-nmea/gps"
	"github.com/twpayne/go-nmea/ublox"
)

var sentenceRx = regexp.MustCompile(`\$[A-Z]+,[^*]+\*(?:[0-9A-Fa-f]{2})?`)

func run() error {
	parser := nmea.NewParser(
		nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
		nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
		nmea.WithSentenceParserFunc(garmin.SentenceParserFunc),
		nmea.WithSentenceParserFunc(gps.SentenceParserFunc),
		nmea.WithSentenceParserFunc(ublox.SentenceParserFunc),
	)

	scanner := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	for scanner.Scan() {
		match := sentenceRx.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		var value any
		switch sentence, err := parser.ParseString(match[0]); {
		case err == nil:
			value = map[string]any{
				sentence.Address().String(): sentence,
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

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
