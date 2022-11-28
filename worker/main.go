package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/michaelChunEC/temporal-helloworld"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort:  "your-custom-namespace.tmprl.cloud:7233",
		Namespace: "your-custom-namespace",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})

	w.RegisterWorkflow(helloworld.Workflow)
	w.RegisterWorkflow(helloworld.Workflow2)
	w.RegisterActivity(helloworld.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
