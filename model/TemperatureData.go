package model

type TemperatureData struct {
	Temperatures []float64
	Stats        Stats
}

type Stats struct {
	Current, Change string
}
