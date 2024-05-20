package asana

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/devzcraft/assignment/internal/config"
)

type Client struct {
	http              *resty.Client
	bucket            chan struct{}
	requestsPerMinute int
}

// TODO: implement stop func to stop ticker
func NewClient(c *config.Config) *Client {
	client := resty.New().
		SetBaseURL(c.Asana.BaseURL).
		SetHeader("Authorization", "Bearer "+c.Asana.Token)

	requestsPerMinte, err := strconv.Atoi(c.RateLimit)
	if err != nil {
		panic("Wrong rate limiter param")
	}

	bucket := make(chan struct{}, requestsPerMinte)

	ticker := time.NewTicker(1 * time.Minute)

	asanaClient := &Client{
		http:              client,
		bucket:            bucket,
		requestsPerMinute: requestsPerMinte,
	}

	go asanaClient.fillBucket(ticker.C)

	return asanaClient
}

func (c *Client) HTTP() *resty.Client {
	return c.http
}

// User retrieve user data from Asana by GID
func (c *Client) User(userGID string) (*resty.Response, error) {
	<-c.bucket
	resp, err := c.http.R().
		Get("users/" + userGID)
	if err != nil {
		return nil, fmt.Errorf("asana client get users error: %w", err)
	}

	return resp, nil
}

// Project retrieve project data from Asana by GID
func (c *Client) Project(projectGID string) (*resty.Response, error) {
	<-c.bucket

	resp, err := c.http.R().
		Get("projects/" + projectGID)
	if err != nil {
		return nil, fmt.Errorf("asana client get projects error: %w", err)
	}

	return resp, nil
}

// fillBucket fill bucket every minute. By default limit is 150 reqeusts per second
func (c *Client) fillBucket(ticker <-chan time.Time) {
	// TODO: if on the first minute bucket is full and we can't send to bucket again
	// it is possible that next minute we will fill bucket with 300 requests
	for i := 0; i < c.requestsPerMinute; i++ {
		c.bucket <- struct{}{}
	}
	for {
		<-ticker
		for i := 0; i < c.requestsPerMinute; i++ {
			c.bucket <- struct{}{}
		}

	}
}
