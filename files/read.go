package files

import (
	"bufio"
	"os"
)

type FnLineProcessor func(line string) (bool, error)

func ReadLines(path string, processor FnLineProcessor) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// TODO: Is the max line size a problem?
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		keepGoing, err := processor(line)
		if err != nil {
			return err
		}
		if !keepGoing {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
