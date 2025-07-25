package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

type GrakeConfig struct {
	Title  string
	Author string
	Files  []string
	Output string
}

const FileName = "grake.toml"

func WriteFile(gptr *GrakeConfig) error {
	var (
		buf []byte
		err error
	)
	if buf, err = toml.Marshal(gptr); err != nil {
		return err
	}
	if err = os.WriteFile(FileName, buf, 0644); err != nil {
		return err
	}
	return nil
}

func ReadFile() (*GrakeConfig, error) {
	var (
		buf  []byte
		err  error
		gptr = new(GrakeConfig)
	)

	if buf, err = os.ReadFile(FileName); err != nil {
		return nil, err
	}
	if err = toml.Unmarshal(buf, gptr); err != nil {
		return nil, err
	}
	return gptr, nil
}
