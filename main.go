package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	// "github.com/Traceableai/goagent"
	// "github.com/Traceableai/goagent/config"
	// "github.com/Traceableai/goagent/instrumentation/net/traceablehttp"

	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {

	// cfg := config.Load()

	// fmt.Println("init of the traceable agent", cfg)
	// shutdown := goagent.Init(cfg)
	// defer shutdown()

	// START OTEL
	ctx := context.Background()

	// Configure a new exporter using environment variables for sending data to Honeycomb over gRPC.
	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	// Create a new tracer provider with a batch span processor and the otlp exporter.
	tp := newTraceProvider(exp)

	// Handle this error in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	// Set the Tracer Provider and the W3C Trace Context propagator as globals
	otel.SetTracerProvider(tp)

	// END OTEL

	router := mux.NewRouter()

	router.Handle("/login", otelhttp.NewHandler(http.HandlerFunc(Login), "/login"))
	router.Handle("/refresh", otelhttp.NewHandler(http.HandlerFunc(Refresh), "/refresh"))
	router.Handle("/test/{id}", otelhttp.NewHandler(isAuthorized(test), "/test/{id}")).Methods("GET")
	router.Handle("/customer/all", otelhttp.NewHandler(isAuthorized(customercount), "/customer/all")).Methods("GET")
	router.Handle("/customer/byid/{id}", otelhttp.NewHandler(isAuthorized(customerbyid), "/customer/byid/{id}")).Methods("GET")
	router.Handle("/crypto/home", otelhttp.NewHandler(isAuthorized(cryptohome), "/crypto/home")).Methods("GET")
	router.Handle("/crypto/price", otelhttp.NewHandler(isAuthorized(cryptoprice), "/crypto/price")).Methods("GET")

	/*
		router.HandleFunc("/login", Login).Methods("GET")
		router.HandleFunc("/refresh", Refresh).Methods("GET")
		router.Handle("/test/{id}", isAuthorized(test)).Methods("GET")
		router.Handle("/customer/all", isAuthorized(customercount)).Methods("GET")
		router.Handle("/customer/byid/{id}", isAuthorized(customerbyid)).Methods("GET")
		router.Handle("/crypto/home", isAuthorized(cryptohome)).Methods("GET")
		router.Handle("/crypto/price", isAuthorized(cryptoprice)).Methods("GET")
	*/
	//Start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint("api.honeycomb.io:443"),
		otlptracegrpc.WithHeaders(map[string]string{
			"x-honeycomb-team":    os.Getenv("HONEYCOMB_API_KEY"),
			"x-honeycomb-dataset": os.Getenv("HONEYCOMB_DATASET"),
		}),
		otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}

	client := otlptracegrpc.NewClient(opts...)
	return otlptrace.New(ctx, client)
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	// The service.name attribute is required.
	resource :=
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ExampleService"),
		)

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)
}
