package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	failureRatio := 0

	go setupMonitor(failureRatio)

	router := gin.Default()
	router.GET("/", thanksHandler)

	router.Run(":3000")
}

// setupMonitor is used to setup a listener in application itself.
// While we make it have a failure ratio which is ratio%(if ratio is 50, then 50%)
// monitor goroutine will listen on port 4000
// Then in HealthCheck of Dockerfile, it should be
// HEALTHCHECK CMD curl –f http://localhost:4000/ || exit 1
func setupMonitor(ratio int) {
	// TODO: add a validation for ratio
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := r.Intn(100)

	fmt.Printf("Got a random number of %d \n", randomNum)

	if randomNum < ratio {
		return
	}

	router := gin.Default()
	router.GET("/", monitorHandler)

	router.Run(":4000")
}

func thanksHandler(c *gin.Context) {
	cmd := exec.Command("cat", "/etc/hostname")

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.String(http.StatusInternalServerError, "Something went wrong :(")
		return
	}

	name := string(output)
	thanks := "Thanks for Attending Docker Meetup!"

	msg := fmt.Sprintf("%s \n\n^^\n I am %s\nfor V1\n\n", thanks, name)

	c.String(http.StatusOK, msg)
}

func monitorHandler(c *gin.Context) {
	c.String(http.StatusOK, "Minitoring Handler is OK ^^")
}
