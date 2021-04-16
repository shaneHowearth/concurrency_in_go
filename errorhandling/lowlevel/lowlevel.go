// Package lowlevel -
package lowlevel

import (
	"os"

	"github.com/shanehowearth/concurrency_in_go/errorhandling/errortype"
)

// Err -
type Err struct {
	error
}

// IsGloballyExec -
func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, Err{(errortype.WrapError(err, err.Error()))}
	}

	return info.Mode().Perm()&0100 == 0100, nil
}
