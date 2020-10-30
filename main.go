package main

import (
	"fmt"
)

type BusinessService interface {
	DoBusiness()
}

type business struct{}

func (b *business) DoBusiness() {
	fmt.Println("Doing Business!")
}

type Tracer struct {
	BusinessService BusinessService
}

func (tracer *Tracer) DoBusiness() {
	fmt.Println("Tracer doing..")
	tracer.BusinessService.DoBusiness()
	fmt.Println("Tracer done!")
}

func NewTracerMiddleware(business BusinessService) *Tracer {
	return &Tracer{
		BusinessService: business,
	}
}

type Metric struct {
	BusinessService BusinessService
}

func (metric *Metric) DoBusiness() {
	fmt.Println("Metric doing...")
	metric.BusinessService.DoBusiness()
	fmt.Println("Metric done!")
}

func NewMetricMiddleware(business BusinessService) *Metric {
	return &Metric{
		BusinessService: business,
	}
}

func main() {
	business := &business{}
	tracer := NewTracerMiddleware(business)
	metric := NewMetricMiddleware(tracer)
	metric.DoBusiness()
}
