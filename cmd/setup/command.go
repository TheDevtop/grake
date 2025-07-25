package setup

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheDevtop/grake/internal/conf"
	"github.com/TheDevtop/grake/internal/defaults"
)

// Compose manuscript file from template
func compose(title, author string, columns uint) string {
	const base = `.TL
%s
.AU
%s
.%dC
.SH
Heading
.LP
Lorem ipsum dolor sit amet.
`
	return fmt.Sprintf(base, title, author, columns)
}

func CmdMain() {
	// Define and parse flags
	var (
		flagTitle  = flag.String("t", defaults.DefaultTitle, "Specify title")
		flagAuthor = flag.String("a", defaults.DefaultAuthor, "Specify name of author")
		flagColumn = flag.Uint("c", defaults.DefaultColumns, "Specify columns")
		flagFile   = flag.String("f", defaults.DefaultSource, "Specify initial source file")
		flagOutput = flag.String("o", defaults.DefaultOutput, "Specify output file")
		flagDir    = flag.String("d", "", "Specify working directory")
	)
	flag.Parse()

	var (
		err  error
		gptr = new(conf.GrakeConfig)
	)

	// Change directory if specified
	if *flagDir != "" {
		if err = os.Chdir(*flagDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Configure GrakeConfig
	gptr.Title = *flagTitle
	gptr.Author = *flagAuthor
	gptr.Files = []string{*flagFile}
	gptr.Output = *flagOutput

	// Write GrakeConfig
	if err = conf.WriteFile(gptr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Compose and write manuscript file
	if err = os.WriteFile(*flagFile, []byte(compose(*flagTitle, *flagAuthor, *flagColumn)), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Initialized: %s\n", gptr.Title)
	os.Exit(0)
}
