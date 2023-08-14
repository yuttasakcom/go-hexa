package middleware

import (
	"context"
	"fmt"

	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/yuttasakcom/go-hexa/src/core/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.GetTracerProvider().Tracer("fiber-server")

const contextLocalKey = "fiber-otel-tracer"

func Tracer(c common.Ctx) error {
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	carrier := propagation.HeaderCarrier{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		carrier.Set(string(key), string(value))
	})

	propagator.Inject(c.Context(), carrier)
	parentCtx := propagator.Extract(context.Background(), carrier)

	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
		trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
		trace.WithAttributes(semconv.NetHostIPKey.String(c.IP())),
		trace.WithAttributes(semconv.HTTPUserAgentKey.String(string(c.Request().Header.UserAgent()))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	ctx, span := tracer.Start(parentCtx, fmt.Sprintf("%s %s", c.Method(), c.Path()), spanOptions...)
	defer span.End()

	c.Locals(contextLocalKey, ctx)
	{
		propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		carrier := propagation.HeaderCarrier{}
		propagator.Inject(ctx, carrier)

		for _, k := range carrier.Keys() {
			c.Response().Header.Set(k, carrier.Get(k))
		}
	}

	err := c.Next()

	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))

	return err
}

func GetSpanContext(c common.Ctx) context.Context {
	ctx, ok := c.Locals(contextLocalKey).(context.Context)
	if !ok {
		slog.L().Warn("Failed to get span context from fiber context")
		return c.Context()
	}

	return ctx
}
