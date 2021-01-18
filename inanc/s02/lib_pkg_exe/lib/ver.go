package lib

import "runtime"

func Ver() string {
	return runtime.Version()
}
