package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/devzcraft/assignment/internal/asana"
	"github.com/devzcraft/assignment/internal/config"
)

func main() {
	conf := config.MustLoad()

	asanaClient := asana.NewClient(conf)

	delay := fetchDelay()

	for {
		<-time.Tick(delay)
		runExtractor(asanaClient)
	}
}

type userProjectExtractor interface {
	projectExtractor
	userExtractor
}

func runExtractor(client userProjectExtractor) {
	users := [4]string{"1207357736997296", "1203393881756367", "1203393881918474", "1207357617619124"}
	for _, gid := range users {
		go extractUser(client, gid)
	}

	projectGID := "1207357733538881"
	go extractProject(client, projectGID)
}

type userExtractor interface {
	User(gid string) (*resty.Response, error)
}

func extractUser(client userExtractor, gid string) {
	user, err := client.User(gid)
	if err != nil {
		// todo log errors
		fmt.Printf("can't get user: %+v", err)
	}

	fmt.Printf("User: %+v\n", user)
}

type projectExtractor interface {
	Project(gid string) (*resty.Response, error)
}

func extractProject(client projectExtractor, gid string) {
	project, err := client.Project(gid)
	if err != nil {
		// todo log errors
		fmt.Printf("can't get project: %+v", err)
	}

	fmt.Printf("Project: %+v\n", project)
}

func fetchDelay() time.Duration {
	//TODO: fetch delay from cli param
	delay, err := time.ParseDuration("30s")
	if err != nil {
		panic("Wrong delay param")
	}

	return delay
}
