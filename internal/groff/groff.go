package groff

import (
	"bytes"
	"os/exec"

	"github.com/TheDevtop/grake/internal/conf"
)

// Format arguments
func argFmt(argv []string) *exec.Cmd {
	var (
		cmd  *exec.Cmd
		args = make([]string, 0, len(argv)+2)
	)

	args = append(args, "-ms")
	args = append(args, argv...)
	args = append(args, "-Tpdf")

	cmd = exec.Command("groff", args...)
	return cmd
}

// Render document
func Render(gptr *conf.GrakeConfig) ([]byte, error) {
	var (
		cmd *exec.Cmd
		err error
		buf = new(bytes.Buffer)
	)

	cmd = argFmt(gptr.Files)
	cmd.Stdout = buf
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
