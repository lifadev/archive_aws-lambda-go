package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var f, h, p string
	var n bool

	flag.StringVar(&f, "function", "handler", "AWS Lambda function name")
	flag.StringVar(&h, "handler", "handle", "AWS Lambda handler name")
	flag.StringVar(&p, "package", "handler.zip", "AWS Lambda package name")
	flag.BoolVar(&n, "nopackage", false, "Only build and do not create AWS Lambda package")

	flag.Parse()

	if len(flag.Args()) > 0 {
		flag.Usage()
	}

	info, err := os.Stat(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	uid := int(info.Sys().(*syscall.Stat_t).Uid)
	gid := int(info.Sys().(*syscall.Stat_t).Gid)

	so := fmt.Sprintf("%s.so", f)

	os.RemoveAll(so)

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-ldflags=-w -s", "-o", so)
	env := os.Environ()
	env = append(env, fmt.Sprintf("CGO_CFLAGS=-DFUNCTION=%s -DHANDLER=%s", f, h))
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chown(so, uid, gid)

	if n {
		return
	}

	os.RemoveAll(p)

	cmd = exec.Command("zip", "-q", "-9", p, so)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chown(p, uid, gid)

	os.RemoveAll(so)
}
