package main

import (
	"dfsa"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	LowerBound = "lowerbound"
	EomLee     = "eomlee"
	Schoute    = "schoute"
	Vogt       = "vogt"
)

type config struct {
	Protocol               string  `json:"protocol"`
	InitialTagsNumber      int     `json:"initialTagsNumber"`
	TagsNumberIncreaseRate float64 `json:"tagsNumberIncreaseRate"`
	MaximumTagsNumber      int     `json:"maximumTagsNumber"`
	Iterations             int     `json:"iterations"`
	InitialFrameSize       int     `json:"initialFrameSize"`
}

type MediaSimulationResult struct {
	TagNumber       int
	SlotsSum        float64
	EmptySlots      float64
	SuccessfulSlots float64
	CollisionSlots  float64
	SimulationTime  float64
	EstimationTime  float64
	EstimationError float64
}

func (c *config) load(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic("Could not open config file in: " + filepath)
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
	flag.StringVar(&configPath, "config", "/home/cilento/lgc/dfsa-simulator/etc/config.json", "Path to configuration file")
	flag.Parse()

	var configuration config
	configuration.load(configPath)

	var estimator dfsa.Estimator

	switch strings.ToLower(configuration.Protocol) {
	case LowerBound:
		estimator = &dfsa.LowerBound{}
	case EomLee:
		estimator = &dfsa.EomLee{}
	case Schoute:
		estimator = &dfsa.Schoute{}
	case Vogt:
		estimator = &dfsa.Vogt{}
	default:
		panic("Invalid Estimator")
	}

	//tagsNumber := configuration.InitialTagsNumber

	var media MediaSimulationResult
	var medias []MediaSimulationResult
	//var quantidadePontos = float32(configuration.MaximumTagsNumber-configuration.InitialTagsNumber) / float32(configuration.TagsNumberIncreaseRate)

	for i := configuration.InitialTagsNumber; i <= configuration.MaximumTagsNumber; i = i + int(configuration.TagsNumberIncreaseRate) {
		//results = nil
		var result dfsa.SimulationResult
		media.TagNumber = 0.0
		media.SlotsSum = 0.0
		media.EmptySlots = 0.0
		media.SuccessfulSlots = 0.0
		media.CollisionSlots = 0.0
		media.SimulationTime = 0.0
		media.EstimationError = 0.0
		var tagsNumber = i
		//fmt.Println(media)
		fmt.Println(i)
		for j := 0; j < configuration.Iterations; j++ {
			simulator := dfsa.Simulator{
				Estimator:        estimator,
				InitialTagsLen:   tagsNumber,
				InitialFrameSize: configuration.InitialFrameSize,
			}

			var simulationTimeInit = time.Now()
			result = simulator.Run()
			var simulationTime = float64(time.Since(simulationTimeInit) / time.Microsecond)

			media.SlotsSum += float64(result.SlotsSum)
			media.EmptySlots += float64(result.EmptySlots)
			media.SuccessfulSlots += float64(result.SuccessfulSlots)
			media.CollisionSlots += float64(result.CollisionSlots)
			media.SimulationTime += simulationTime
			media.EstimationTime += float64(result.EstimationTime)
			media.EstimationError += result.EstimationError
		}

		media.TagNumber = tagsNumber
		media.SlotsSum = media.SlotsSum / float64(configuration.Iterations)
		media.EmptySlots = media.EmptySlots / float64(configuration.Iterations)
		media.SuccessfulSlots = media.SuccessfulSlots / float64(configuration.Iterations)
		media.CollisionSlots = media.CollisionSlots / float64(configuration.Iterations)
		media.SimulationTime = media.SimulationTime / float64(configuration.Iterations)
		media.EstimationTime = media.EstimationTime / float64(configuration.Iterations)
		media.EstimationError = media.EstimationError / float64(configuration.Iterations)

		medias = append(medias, media) //fmt.Println(media)
	}
	file, err := os.Create("result_" + strings.ToLower(configuration.Protocol) + ".csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//var header []string
	//header = []string{"TagNumber", "SlotsSum", "EmptySlots", "SuccessfulSlots", "CollisionSlots", "SimulationTime", "EstimationTime", "EstimationError"}
	//err = writer.Write(header)
	//checkError("Cannot write to file", err)
	for _, value := range medias {
		str1 := strconv.Itoa(value.TagNumber)
		str2 := strconv.Itoa(int(math.Ceil(value.SlotsSum)))
		str3 := strconv.Itoa(int(math.Ceil(value.EmptySlots)))
		str4 := strconv.Itoa(int(math.Ceil(value.SuccessfulSlots)))
		str5 := strconv.Itoa(int(math.Ceil(value.CollisionSlots)))
		str6 := strconv.Itoa(int(value.SimulationTime))
		str7 := strconv.Itoa(int(value.EstimationTime))
		str8 := strconv.Itoa(int(value.EstimationError))
		var str []string
		str = append(str, str1)
		str = append(str, str2)
		str = append(str, str3)
		str = append(str, str4)
		str = append(str, str5)
		str = append(str, str6)
		str = append(str, str7)
		str = append(str, str8)
		err := writer.Write(str)
		checkError("Cannot write to file", err)
	}

	//fmt.Println(medias)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
