package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gustavolopess/dfsa-simulator/src/dfsa"
	"io/ioutil"
	"os"
	"strings"
)

const (
	LowerBound = "lowerbound"
	EomLee = "eomlee"
)

type config struct {
	Protocol string `json:"protocol"`
	InitialTagsNumber int `json:"initialTagsNumber"`
	TagsNumberIncreaseRate float64 `json:"tagsNumberIncreaseRate"`
	MaximumTagsNumber int `json:"maximumTagsNumber"`
	Iterations int `json:"iterations"`
	InitialFrameSize int `json:"initialFrameSize"`
}

func (c *config) load(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic("Could not open config file")
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Could not read config file")
	}
	err = json.Unmarshal(fileBytes, c)
	if err != nil {
		panic("Invalid config file JSON")
	}
}


func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "etc/config.json", "Path to configuration file")
	flag.Parse()

	var configuration config
	configuration.load(configPath)

	var estimator dfsa.Estimator

	switch strings.ToLower(configuration.Protocol) {
	case LowerBound:
		estimator = &dfsa.LowerBound{}
	case EomLee:
		estimator = &dfsa.EomLee{}
	default:
		panic("Invalid Estimator")
	}

	tagsNumber := configuration.InitialTagsNumber
	for i := 0; i < configuration.Iterations; i++ {
		simulator := dfsa.Simulator{
			Estimator: estimator,
			InitialTagsLen: tagsNumber,
			InitialFrameSize: configuration.InitialFrameSize,
		}
		fmt.Println(simulator.Run())
	}
}
