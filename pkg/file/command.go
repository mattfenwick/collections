package file

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd) (string, error) {
	if len(cmd.Dir) == 0 {
		dir, _ := os.Getwd()
		logrus.Debugf("running command: '%s' in current directory '%s'", cmd.String(), dir)
	} else {
		logrus.Debugf("running command: '%s' in directory '%s'", cmd.String(), cmd.Dir)
	}
	cmdOutput, err := cmd.CombinedOutput()
	cmdOutputStr := string(cmdOutput)
	logrus.Tracef("command: '%s' output:\n%s", cmd.String(), cmdOutput)
	return cmdOutputStr, errors.Wrapf(err, "unable to run command '%s': %s", cmd.String(), cmdOutputStr)
}

func ExtendCommandEnvironment(cmd *exec.Cmd, env map[string]string) {
	cmd.Env = os.Environ()
	for key, val := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, val))
	}
}

// RunCommandWithPipes runs a long running command and continuously streams its output
func RunCommandWithPipes(cmd *exec.Cmd) error {
	// without this, stdin is constantly fed input
	_, _ = cmd.StdinPipe()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// can't use RunCommand(cmd) here -- attaching to os pipes interferes with cmd.CombinedOutput()
	logrus.Infof("running command '%s' with pipes attached in directory '%s' and with env \n%+v\n", cmd.String(), cmd.Dir, cmd.Env)
	return errors.Wrapf(cmd.Run(), "unable to run command '%s'", cmd.String())
}
