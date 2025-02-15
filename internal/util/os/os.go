package os

import (
	"os"
	"strings"

	"github.com/mickael-menu/zk/internal/util/opt"
)

// Getenv returns an optional String for the environment variable with given
// key.
func GetOptEnv(key string) opt.String {
	return opt.NewNotEmptyString(os.Getenv(key))
}

// Env returns a map of environment variables.
func Env() map[string]string {
	env := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		env[pair[0]] = pair[1]
	}
	return env
}
