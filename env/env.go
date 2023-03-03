package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/j13g/goutil/types"
	"github.com/samber/mo"
)

// allow setting a global environment prefix
var envPrefix = ""

func SetEnvPrefix(prefix string) {
	envPrefix = prefix
}

type Option func(options *envOptions)

type envOptions struct {
	usePrefix string
}

func (eo envOptions) envName(name string) string {
	if eo.usePrefix == "" {
		return strings.ToUpper(name)
	}
	return fmt.Sprintf("%s_%s", strings.ToUpper(eo.usePrefix), strings.ToUpper(name))
}

func WithPrefix(prefix string) Option {
	return func(options *envOptions) {
		options.usePrefix = prefix
	}
}

func handleOptions(options []Option) *envOptions {
	opt := &envOptions{
		usePrefix: envPrefix,
	}
	for _, o := range options {
		o(opt)
	}
	return opt
}

func GetString(key string, options ...Option) mo.Option[string] {
	o := handleOptions(options)
	envName := o.envName(key)
	val, ok := os.LookupEnv(envName)
	if ok {
		return mo.Some(val)
	}
	return mo.None[string]()
}

func GetStringDefault(key string, defaultValue string, options ...Option) string {
	return GetString(key, options...).OrElse(defaultValue)
}

func GetInt(key string) mo.Option[int] {
	val := GetString(key)
	if val.IsAbsent() {
		return mo.None[int]()
	}

	i, err := strconv.Atoi(val.MustGet())
	if err != nil {
		return mo.None[int]()
	}
	return mo.Some(i)
}

func GetIntDefault(key string, defaultValue int) int {
	return GetInt(key).OrElse(defaultValue)
}

func GetFloat(key string) mo.Option[float64] {
	val := GetString(key)
	if val.IsAbsent() {
		return mo.None[float64]()
	}
	f, err := strconv.ParseFloat(val.MustGet(), 64)
	if err != nil {
		return mo.None[float64]()
	}

	return mo.Some(f)
}

func GetFloatDefault(key string, defaultValue float64) float64 {
	return GetFloat(key).OrElse(defaultValue)
}

var trueValues = types.NewSetFromSlice[string]([]string{
	"TRUE",
	"T",
	"YES",
})

var falseValues = types.NewSetFromSlice[string]([]string{
	"FALSE",
	"F",
	"NO",
})

func GetBool(key string) mo.Option[bool] {
	x := GetString(key)
	if x.IsAbsent() {
		return mo.None[bool]()
	}

	val := x.MustGet()
	val = strings.ToUpper(val)

	if trueValues.Contains(val) {
		return mo.Some(true)
	}
	if falseValues.Contains(val) {
		return mo.Some(false)
	}
	return mo.None[bool]()
}
