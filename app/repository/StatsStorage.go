package repository

import (
	"log"
	"math"
	"os"
)

type StatResults struct {
	Total   int     `json:"total"`
	Average float64 `json:"average"`
}

type Stats struct {
	logger      *log.Logger
	postRequest []int64
	average     float64
}

func NewStats() *Stats {
	logger := log.New(os.Stdout, "StatsStorage Repository: ", log.LstdFlags)
	postRequest := []int64{}
	average := 0.0
	return &Stats{logger, postRequest, average}
}

func (repository *Stats) AddStat(time int64) {

	repository.postRequest = append(repository.postRequest, time)

	var timeSum = int64(0)
	for _, postTime := range repository.postRequest {
		timeSum += postTime
	}

	repository.average = math.Round((float64(timeSum)/float64(len(repository.postRequest)))*100) / 100
}

func (repository *Stats) GetStat() StatResults {
	return StatResults{len(repository.postRequest), repository.average}
}
