package main

import (
	"bufio"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/kdar/tvbgonegen"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.App("tvbgonegen", "TV-B-Gone Gen")
	app.Spec = "[--freq] [--struct-name] [--margin-of-error] [SRC]"
	var (
		structName    = app.StringOpt("struct-name", "code", "the name of the code structure")
		frequency     = app.IntOpt("freq", 38462, "the frequency of the code")
		src           = app.StringArg("SRC", "-", "the file to parse. defaults to '-' for STDIN")
		marginOfError = app.IntOpt("margin-of-error", 12, "the margin of error that is acceptable before two timings are different")
	)
	app.Action = func() {
		var err error
		var input []byte
		if *src == "-" {
			if terminal.IsTerminal(int(os.Stdin.Fd())) {
				in := bufio.NewReader(os.Stdin)
				input, err = in.ReadSlice(0x04)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				input, err = ioutil.ReadAll(os.Stdin)
			}
		} else {
			input, err = ioutil.ReadFile(*src)
		}
		if err != nil {
			log.Fatal(err)
		}

		gen := tvbgonegen.New()
		gen.MarginOfError = *marginOfError
		gen.TimesName = *structName + "_times"
		gen.CodesName = *structName
		gen.Frequency = *frequency
		data, err := gen.IRremote(tvbgonegen.ParseIRremote(input))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data.Format(tvbgonegen.DefaultTemplate))
	}
	app.Run(os.Args)
}
