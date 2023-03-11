package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	syscall.Chroot("/home/containers/makeyourowncontainer/alpine-minirootfs-3.17.2-aarch64")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	cmd := exec.Command("/bin/busybox", "sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stdout = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      1000,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      1000,
				Size:        1,
			},
		},
	}

	syscall.Unmount("proc", 0)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
