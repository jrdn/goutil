package cli

import (
	"errors"
	"io"
	"strings"

	"github.com/jrdn/goutil/cli/outputter"
)

type Outputter interface {
	Format(x any) (string, error)
	Print(x any) error
	Write(w io.Writer, x any) error
}

var globalOutputter Outputter = outputter.JSONOutputter{}

func SetOutputter(o Outputter) {
	globalOutputter = o
}

func SetOutputterByName(s string) error {
	switch strings.ToLower(s) {
	case "json":
		SetOutputter(outputter.JSONOutputter{})
	case "go":
		SetOutputter(outputter.GoOutputter{})
	default:
		return errors.New("invalid outputter: " + s)
	}
	return nil
}

func Format(x any) (string, error) {
	return globalOutputter.Format(x)
}

func Print(x any) error {
	return globalOutputter.Print(x)
}

func Write(w io.Writer, x any) error {
	return globalOutputter.Write(w, x)
}
