package wizard

import (
	"errors"
	"io"
	"os"
	"strings"
)

func ReadStdinLine() (string, error) {
	return readLine(os.Stdin)
}

func readLine(r io.Reader) (string, error) {
	output := strings.Builder{}

	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if buf[0] == '\n' {
			break
		}
		err = output.WriteByte(buf[0])
		if err != nil {
			return "", err
		}
	}

	return output.String(), nil
}
