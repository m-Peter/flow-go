package main

import (
	"flag"
	"os"

	"github.com/onflow/flow-go/integration/benchnet2/automate/level1"
)

// sample usage:
// go run cmd/level1/bootstrap.go  --data "./testdata/level1/data/root-protocol-state-snapshot1.json" --docker-tag "v0.27.6"
func main() {
	dataFlag := flag.String("data", "", "Path to bootstrap JSON data.")
	dockerTagFlag := flag.String("docker-tag", "", "Docker image tag.")
	flag.Parse()

	if *dataFlag == "" || *dockerTagFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	bootstrap := level1.NewBootstrap(*dataFlag)
	bootstrap.GenTemplateData(true, *dockerTagFlag)
}
