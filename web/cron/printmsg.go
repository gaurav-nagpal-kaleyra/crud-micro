package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

// 2
func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func RunCronJobs() {
	// 3
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	_, err := client.Head("https://hc-ping.com/0c948df7-30be-456c-a5d3-5daf4afd4283")
	if err != nil {
		fmt.Printf("error making start request%s", err)
	}
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(5).Minutes().Do(func() {
		hello("Mac")
	})
	_, err = client.Head("https://hc-ping.com/0c948df7-30be-456c-a5d3-5daf4afd4283")
	if err != nil {
		fmt.Printf("error making sucess request%s", err)
	}
	// 5
	s.StartBlocking()

}
