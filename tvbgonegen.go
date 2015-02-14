package tvbgonegen

import (
	"bytes"
	"fmt"
	"math"
	"text/template"
)

//"strconv"

var (
	DefaultTemplate = `const uint16_t {{.TimesName}}[] PROGMEM = {{"{"}}{{range $_, $v := .Times}}
  {{$v}},{{end}}
};

const struct IrCode {{.CodesName}} PROGMEM = {
  freq_to_timerval({{.Frequency}}),
  {{.CodePairs}},
  {{.BitLength}},
  {{.TimesName}},
  {{"{"}}{{range $_, $v := .Codes}}
    0x{{intToHex $v}},{{end}}
  }
};

`
)

type Gen struct {
	MarginOfError int
	TimesName     string
	CodesName     string
	Frequency     int
}

func New() *Gen {
	return &Gen{}
}

type Data struct {
	*Gen
	CodePairs int
	Times     []int
	Codes     []uint8
	BitLength int
}

func (d *Data) Format(tpl string) (string, error) {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"intToHex": func(n interface{}) string {
			return fmt.Sprintf("%X", n)
		},
		"divide": func(a, b int) int {
			return a / b
		},
	}

	tmpl, err := template.New("tpl").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return "", fmt.Errorf("parsing: %s", err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, d)
	if err != nil {
		return "", fmt.Errorf("execution: %s", err)
	}

	return buf.String(), nil
}

func BitLength(n int) int {
	// Determine how many bits we need to use
	bitLen := 0
	for i := 8; i >= 0; i-- {
		if (1<<uint(i-1))&(n) > 0 {
			bitLen = i
			break
		}
	}
	return bitLen
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func absu16(n uint16) uint16 {
	return uint16(math.Abs(float64(n)))
}
