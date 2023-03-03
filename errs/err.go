package errs

import "fmt"

func Wrap(base error, msg string) error {
	return fmt.Errorf("%w: %s", base, msg)
}
