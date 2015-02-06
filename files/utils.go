package files

import (
	"os"

	"github.com/justinsb/gova/errors"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ListFilenames(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, errors.Chain(err, "Error opening directory ", dir)
	}

	defer f.Close()

	return f.Readdirnames(-1)
}
