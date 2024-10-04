package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	images, err := apiClient.ImageList(context.Background(), image.ListOptions{All: true})

	_, inspectJson, err := apiClient.ImageInspectWithRaw(context.Background(), images[0].ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", inspectJson)
}
