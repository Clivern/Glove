// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"math/rand"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

// Exporter type
type Exporter struct {
	mutex sync.Mutex
	up    *prometheus.Desc
}

// NewExporter creates a new Exporter
func NewExporter() *Exporter {
	return &Exporter{
		up: prometheus.NewDesc(
			prometheus.BuildFQName(viper.GetString("app.name"), "", "up"),
			"Could service be reached?",
			[]string{"name", "uri"},
			nil),
	}
}

// Collect and Sends the collected metrics to Prometheus
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	ch <- prometheus.MustNewConstMetric(
		e.up,
		prometheus.GaugeValue,
		rand.Float64(),
		"example",
		"http://example.com",
	)
}

// Describe exposes the metric description to Prometheus
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
}
