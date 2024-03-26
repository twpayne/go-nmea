package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/gps"
	"github.com/twpayne/go-nmea/nmeapgrm"
	"github.com/twpayne/go-nmea/nmeapubx"
)

var sentenceRx = regexp.MustCompile(`\$[^*]+\*(?:[0-9A-Fa-f]{2})?`)

func run() error {
	parser := nmea.NewParser(
		nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
		nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
		nmea.WithSentenceParserFunc(gps.SentenceParser),
		nmea.WithSentenceParserFunc(nmeapgrm.SentenceParser),
		nmea.WithSentenceParserFunc(nmeapubx.SentenceParser),
	)

	scanner := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	for scanner.Scan() {
		match := sentenceRx.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		sentence, err := parser.ParseString(match[0])
		if err != nil {
			continue // FIXME
		}
		if err := encoder.Encode(map[string]any{
			sentence.Address().String(): sentence,
		}); err != nil {
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
