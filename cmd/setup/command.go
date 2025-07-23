package setup

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheDevtop/grake/internal/conf"
)

// Compose manuscript file from template
func compose(gptr *conf.GrakeConfig) string {
	const base = `.TL
%s
.AU
%s
.2C
.SH
Heading
.LP
Lorem ipsum dolor sit amet.
`
	return fmt.Sprintf(base, gptr.Title, gptr.Author)
}

func CmdMain() {
	// Define and parse flags
	var (
		flagTitle  = flag.String("t", "Title", "Specify title")
		flagAuthor = flag.String("a", "Author Name", "Specify name of author")
		flagFile   = flag.String("f", "main.ms", "Specify main file")
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

	// Write GrakeConfig
	if err = conf.WriteFile(gptr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Compose and write manuscript file
	if err = os.WriteFile(*flagFile, []byte(compose(gptr)), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
