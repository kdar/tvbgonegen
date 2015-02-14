package tvbgonegen

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

// Decoded NEC: 1CE348B7 (32 bits)
// Raw (68): -31588 7600 -3600 550 -350 600 -350 600 -350 600 -1250 600 -1250 600 -1250 550 -400 550 -350 650 -1200 650 -1200 650 -1200 600 -350 600 -350 550 -350 650 -1200 650 -1200 600 -350 650 -1200 600 -350 550 -400 600 -1250 550 -350 600 -350 600 -350 600 -1250 550 -400 600 -1250 600 -1250 550 -350 650 -1200 650 -1200 650 -1200 600

var (
	ircodesRegex = regexp.MustCompile(`(-?(\d+) -?(\d+) ?)+`)
)

func ParseIRremote(input []byte) []int {
	var timings []int
	match := ircodesRegex.Find(input)
	buf := bytes.NewBuffer(match)
	var d int
	for {
		_, err := fmt.Fscan(buf, &d)
		if err != nil {
			break
		}

		timings = append(timings, d)
	}
	return timings
}

// https://github.com/shirriff/Arduino-IRremote
func (g *Gen) IRremote(timings []int) (*Data, error) {
	var times []int
	var codeTimeIndex []byte

	if len(timings)%2 == 1 {
		return nil, fmt.Errorf("timings must be in groups of two (even). found: %d", len(timings))
	}

	for i := 1; i < len(timings); i += 2 {
		on := abs(timings[i]) / 10
		off := 0
		if i+1 < len(timings) {
			off = abs(timings[i+1]) / 10
		}

		found := false
		timeIndex := 0
		for x := 0; x < len(times); x += 2 {
			found = times[x]-on < g.MarginOfError
			found = found && (abs(times[x+1]-off) < g.MarginOfError || off == 0)
			if found {
				timeIndex = x / 2
				break
			}
		}

		if !found {
			times = append(times, on, off)
			timeIndex = len(times)/2 - 1
		}

		codeTimeIndex = append(codeTimeIndex, byte(timeIndex))
	}

	bitLen := BitLength(len(times)/2 - 1)

	var codesStr string
	for _, v := range codeTimeIndex {
		binary := strconv.FormatInt(int64(v), 2)
		zeros := bitLen - len(binary)
		for i := 0; i < zeros; i++ {
			binary = "0" + binary
		}

		codesStr += binary
	}

	zeros := len(codesStr) % 8
	for i := 0; i < zeros; i++ {
		codesStr = codesStr + "0"
	}

	codebig := big.NewInt(0)
	codebig.SetString(codesStr, 2)

	return &Data{
		Gen:       g,
		CodePairs: len(timings) / 2,
		Times:     times,
		Codes:     codebig.Bytes(),
		BitLength: bitLen,
	}, nil
}
