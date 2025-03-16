package main

import (
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {
	// Initialize Jaeger Exporter
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		log.Fatal(err)
	}

	// Create Trace Provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("app-one"),
		)),
	)

	otel.SetTracerProvider(tp)

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(SimpleHandler), "Hello"))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, World!"))
}
