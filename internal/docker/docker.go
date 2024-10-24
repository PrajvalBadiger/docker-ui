package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type DockerImage struct {
	ID      string
	Repo    string
	Tag     string
	Size    int64
	Created int64
}

type DockerWrapper struct {
	client *client.Client
}

func (c *DockerWrapper) NewClient() {
	var err error
	c.client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err.Error())
	}
}

func (c *DockerWrapper) CloseClient() {
	c.client.Close()
}

func (c *DockerWrapper) GetImages() []DockerImage {
	images, err := c.client.ImageList(context.Background(), image.ListOptions{All: true})

	if err != nil {
		panic(err)
	}

	var dockerImages []DockerImage
	for _, img := range images {
		if len(img.RepoTags) == 0 {
			dockerImages = append(dockerImages, DockerImage{
				ID:      img.ID,
				Repo:    "<none>",
				Tag:     "<none>",
				Size:    img.Size,
				Created: img.Created,
			})
		} else {
			for _, tag := range img.RepoTags {
				repo := strings.Split(tag, ":")
				dockerImages = append(dockerImages, DockerImage{
					ID:      img.ID,
					Repo:    repo[0],
					Tag:     repo[1],
					Size:    img.Size,
					Created: img.Created,
				})
			}
		}
	}

	return dockerImages
}

func (dw *DockerWrapper) ListImages() []string {
	dockerImages := dw.GetImages()

	s := make([]string, 0)
	for _, img := range dockerImages {
		i := img.Repo + ":" + img.Tag
		s = append(s, i)
	}
	return s
}
