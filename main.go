package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

const meterName = "github.com/pedrobarco/otel-tally-counter"
const counterName = "counter"

const serverAddress = ":8080"

func main() {
	ctx := context.Background()

	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(meterName)

	counterMetric, err := meter.Int64Counter(counterName)
	if err != nil {
		log.Fatal(err)
	}

	counter := &CustomCounter{}

	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc(1)
		counterMetric.Add(ctx, 1)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("Counter value: %d", counter.Value())))
		if err != nil {
			log.Fatal(err)
		}
	})

	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Server running on: %s\n", serverAddress)
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
