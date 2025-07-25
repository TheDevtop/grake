package clean

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheDevtop/grake/internal/conf"
)

func CmdMain() {
	var flagDir = flag.String("d", "", "Specify working directory")
	flag.Parse()

	var (
		err  error
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

	if err = os.Remove(gptr.Output); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Cleaned: %s\n", gptr.Title)
	os.Exit(0)
}
