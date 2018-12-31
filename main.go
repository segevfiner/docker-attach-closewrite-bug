package main

import (
	"context"
	"flag"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var closeWrite = flag.Bool("close-write", false, "Call CloseWrite")

func main() {
	flag.Parse()

	ctx := context.Background()

	cli, err := docker.NewClientWithOpts(docker.WithVersion("1.39"), docker.FromEnv)
	if err != nil {
		panic(err)
	}

	containerConfig := &container.Config{
		Image:     "docker-attach-closewrite",
		OpenStdin: true,
	}

	hostConfig := &container.HostConfig{
		AutoRemove: true,
	}

	containerRes, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, "")
	if err != nil {
		panic(err)
	}

	stream, err := cli.ContainerAttach(ctx, containerRes.ID, types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, containerRes.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	stream.Conn.Write([]byte("World\n"))

	if *closeWrite {
		stream.CloseWrite()
	}

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, stream.Reader)
	if err != nil {
		panic(err)
	}
}
