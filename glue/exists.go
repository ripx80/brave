package glue

import (
	"os/exec"
)

/*CmdExist check if the given cmd is available in system path*/
func CmdExist(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
