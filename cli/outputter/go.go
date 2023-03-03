package outputter

import (
	"fmt"
	"io"
)

type GoOutputter struct{}

func (g GoOutputter) Format(x any) (string, error) {
	out := fmt.Sprintf("%+v", x)
	return out, nil
}

func (g GoOutputter) Print(x any) error {
	out, err := g.Format(x)
	fmt.Print(out)
	return err
}

func (g GoOutputter) Write(w io.Writer, x any) error {
	out, err := g.Format(x)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, out)
	return err
}
