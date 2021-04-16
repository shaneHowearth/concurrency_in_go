// Package intermediate -
package intermediate

import (
	"os/exec"

	"github.com/shanehowearth/concurrency_in_go/errorhandling/errortype"
	"github.com/shanehowearth/concurrency_in_go/errorhandling/lowlevel"
)

// Err -
type Err struct {
	error
}

// RunJob -
func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := lowlevel.IsGloballyExec(jobBinPath)
	if err != nil {
		return err
	} else if !isExecutable {
		return errortype.WrapError(nil, "job binary is not executable")
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}
