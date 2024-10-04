package main

import (
	"fmt"

	"github.com/PrajvalBadiger/docker-ui/internal/docker"
)

func main() {
	var dw docker.DockerWrapper

	dw.NewClient()
	defer dw.CloseClient()

	dockerImages := dw.GetImages()

	for _, img := range dockerImages {
		fmt.Println(img.Tag, img.ID, img.Created)
	}
}
