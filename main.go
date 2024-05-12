package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func ContainerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	cli.NegotiateAPIVersion(ctx)
	var dockerresponses []dockerapi
	for _, container := range containers {
		dockerresponse := dockerapi{

			ID:     container.ID,
			Name:   container.Names[0],
			Image:  container.Image,
			State:  container.State,
			Status: container.Status,
			Ports:  container.Ports,
		}
		dockerresponses = append(dockerresponses, dockerresponse)
	}
	jsonResponse, err := json.Marshal(dockerresponses)
	if err != nil {

		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

type dockerapi struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Image  string       `json:"image"`
	State  string       `json:"state"`
	Status string       `json:"status"`
	Ports  []types.Port `json:"ports"`
}

func customHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static/dashboard.html")
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	getimages, err := cli.ImagePull(ctx, "docker.io/nginx:latest", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, getimages)
	getimages.Close()

	containerreouser, err := cli.ContainerCreate(ctx, &container.Config{Image: "nginx"}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}
	if err = cli.ContainerStart(ctx, containerreouser.ID, container.StartOptions{}); err != nil {
		panic(err)

	}
	// Register route handlers
	http.HandleFunc("/ContainerHealth", ContainerHandler)
	http.HandleFunc("/Dashboard", customHandler)

	// Start the HTTP server on port 8080
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
