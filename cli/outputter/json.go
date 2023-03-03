package outputter

import (
	"encoding/json"
	"fmt"
	"io"
)

type JSONOutputter struct{}

func (j JSONOutputter) Format(x any) (string, error) {
	bytes, err := json.MarshalIndent(x, "", "  ")
	return string(bytes), err
}

func (j JSONOutputter) Print(x any) error {
	out, err := j.Format(x)
	fmt.Print(out)
	return err
}

func (j JSONOutputter) Write(w io.Writer, x any) error {
	out, err := j.Format(x)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, out)
	return err
}
