/*
Copyright 2021 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"bailuoxi66/go-loggie/pkg/core/signals"
	"flag"
	"os"
)

var (
	globalConfigFile   string
	pipelineConfigPath string
	nodeName           string
)

func init() {
	hostName, _ := os.Hostname()
	flag.StringVar(&globalConfigFile, "config.system", "loggie.yml", "global config file")
	flag.StringVar(&pipelineConfigPath, "config.pipeline", "pipelines.yml", "reloadable config file path")
	flag.StringVar(&nodeName, "meta.nodeName", hostName, "override nodeName")
}

func main() {
	flag.Parse()
	log.InitLog()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
}
