package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func initTracer() (func(), error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.Default()),
	)
	otel.SetTracerProvider(tp)

	return func() {
		tp.Shutdown()
	}, nil
}

func main() {
	shutdown, err := initTracer()
	if err != nil {
		panic(err)
	}
	defer shutdown()

	r := chi.NewRouter()

	// Routes
	r.Get("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	}), "HomePage").ServeHTTP)

	r.Get("/api/profile", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the profile page!"))
	}), "ProfilePage").ServeHTTP)

	r.Get("/api/profile/{username}", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		w.Write([]byte("Username: " + username))
	}), "UserProfilePage").ServeHTTP)

	// Start the server
	http.ListenAndServe(":8080", otelhttp.NewHandler(r, "Server"))
}