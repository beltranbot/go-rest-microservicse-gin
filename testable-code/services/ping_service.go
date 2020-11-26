package services

import "fmt"

const (
	pong = "pong"
)

type pingService interface {
	HandlePing() (string, error)
}

type pingServiceImpl struct{}

var (
	// PingService public instance
	PingService pingService = pingServiceImpl{}
)

// HandlePing func
func (service pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("doing some complex things...")
	return pong, nil
}
