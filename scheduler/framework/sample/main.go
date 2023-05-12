package main

import (
	"fmt"
	"math/rand"
	"os"
	"my-scheduler/pkg"
	"time"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("scheduler start")
	command := app.NewSchedulerCommand(
		app.WithPlugin(pkg.Name, pkg.New),
	)
	
	fmt.Println("scheduler running")

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Println("scheduler ends")

}
