package build

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheDevtop/grake/internal/conf"
	"github.com/TheDevtop/grake/internal/groff"
)

func CmdMain() {
	var flagDir = flag.String("d", "", "Specify working directory")
	flag.Parse()

	var (
		err  error
		buf  []byte
		gptr *conf.GrakeConfig
	)

	// Change directory if specified
	if *flagDir != "" {
		if err = os.Chdir(*flagDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if gptr, err = conf.ReadFile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Check if all source files are accounted for
	for _, file := range gptr.Files {
		if _, err = os.Stat(file); err != nil {
			fmt.Fprintf(os.Stderr, "Could not find file (%s)\n", err)
			os.Exit(1)
		}
	}

	// Call Groff to render the document
	if buf, err = groff.Render(gptr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Write the buffer to file
	if err = os.WriteFile(gptr.Output, buf, 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Build: %s\n", gptr.Title)
	os.Exit(0)
}
