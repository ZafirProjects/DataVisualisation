package handler

import (
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/ZafirProjects/QuodOrbisChallenge/model"
	"github.com/ZafirProjects/QuodOrbisChallenge/view/generate"
)

type GenerateHandler struct{}

func (h *GenerateHandler) RenderGeneratedData(w http.ResponseWriter, r *http.Request) {
	data := generateTemperatureData(100, 0.4, 0.4)
	render(w, r, generate.RenderGenerateData(data))
}

func (h *GenerateHandler) RenderNewData(w http.ResponseWriter, r *http.Request) {
	data := generateTemperatureData(100, 0.4, 0.4)
	render(w, r, generate.Data(data))
}

func generateTemperatureData(length int, trendStrength float64, noiseStrength float64) model.TemperatureData {
	// Initialize the time series with a random starting point
	series := []float64{math.Round((rand.Float64()+20)*10) / 10}

	// Generate the trend component
	trend := make([]float64, length)
	for i := range trend {
		trend[i] = rand.NormFloat64() * trendStrength
	}

	// Generate the noise component
	noise := make([]float64, length)
	for i := range noise {
		noise[i] = rand.NormFloat64() * noiseStrength
	}

	// Generate the time series
	for i := 1; i < length; i++ {
		// the i-1 is the previous data point hence relative
		temperature := series[i-1] + trend[i] + noise[i]
		temperature = math.Round(temperature*10) / 10
		series = append(series, temperature)
	}

	changeFloat := series[len(series)-1] - series[0]
	changeFloat = math.Round(changeFloat*10) / 10

	return model.TemperatureData{
		Temperatures: series,
		Stats: model.Stats{
			Current: strconv.FormatFloat(series[len(series)-1], 'f', -1, 64),
			Change:  strconv.FormatFloat(changeFloat, 'f', -1, 64),
		},
	}
}
