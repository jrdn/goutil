package outputter

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type YAMLOutputter struct{}

func (y YAMLOutputter) Format(x any) (string, error) {
	out, err := yaml.Marshal(x)
	return string(out), err
}

func (y YAMLOutputter) Print(x any) error {
	out, err := y.Format(x)
	fmt.Println(out)
	return err
}

func (y YAMLOutputter) Write(w io.Writer, x any) error {
	out, err := y.Format(x)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, out)
	return err
}
