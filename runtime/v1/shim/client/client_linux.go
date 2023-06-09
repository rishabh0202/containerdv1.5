//go:build linux
// +build linux

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package client

import (
	"os/exec"
	"syscall"

	"github.com/containerd/cgroups"
	"github.com/pkg/errors"
)

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func setCgroup(cgroupPath string, cmd *exec.Cmd) error {
	cg, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(cgroupPath))
	if err != nil {
		return errors.Wrapf(err, "failed to load cgroup %s", cgroupPath)
	}
	if err := cg.Add(cgroups.Process{
		Pid: cmd.Process.Pid,
	}); err != nil {
		return errors.Wrapf(err, "failed to join cgroup %s", cgroupPath)
	}
	return nil
}
