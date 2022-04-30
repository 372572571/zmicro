package zmicro

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/iobrother/zmicro/core/config"
	"github.com/iobrother/zmicro/core/log"
	"github.com/iobrother/zmicro/core/transport/http"
	"github.com/iobrother/zmicro/core/transport/rpc/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config", "config.yaml", "config file")
}

type App struct {
	opts       Options
	zc         *zconfig
	rpcServer  *server.Server
	httpServer *http.Server
}

type zconfig struct {
	App struct {
		Name           string
		Addr           string
		JaegerEndpoint string
	}
	Http struct {
		Mode string
	}
	Rpc struct {
		BasePath       string
		UpdateInterval int
		EtcdAddr       []string
	}
}

func New(opts ...Option) *App {
	options := newOptions(opts...)
	flag.Parse()
	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		log.Fatal("config file not exists")
	}

	c := config.New(config.Path(cfgFile))
	config.ResetDefault(c)

	zc := &zconfig{}
	if err = config.Scan("app", &zc.App); err != nil {
		log.Fatal(err.Error())
	}
	if err = config.Scan("http", &zc.Http); err != nil {
		log.Fatal(err.Error())
	}
	if err = config.Scan("rpc", &zc.Rpc); err != nil {
		log.Fatal(err.Error())
	}

	app := &App{
		opts: options,
		zc:   zc,
	}

	tracing := false
	if zc.App.JaegerEndpoint != "" {
		setTracerProvider(zc.App.JaegerEndpoint, zc.App.Name)
		tracing = true
	}

	if app.opts.InitRpcServer != nil {
		app.rpcServer = server.NewServer(
			server.Name(zc.App.Name),
			server.BasePath(zc.Rpc.BasePath),
			server.UpdateInterval(zc.Rpc.UpdateInterval),
			server.EtcdAddr(zc.Rpc.EtcdAddr),
			server.Tracing(tracing),
		)
		app.rpcServer.Init(server.InitRpcServer(app.opts.InitRpcServer))
	}
	if app.opts.InitHttpServer != nil {
		app.httpServer = http.NewServer(
			http.Name(zc.App.Name),
			http.Mode(zc.Http.Mode),
			http.Tracing(tracing),
		)
		app.httpServer.Init(http.InitHttpServer(app.opts.InitHttpServer))
	}

	return app
}

func setTracerProvider(endpoint string, name string) *trace.TracerProvider {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		log.Fatal(err.Error())
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp
}

func (a *App) Run() error {

	l, err := net.Listen("tcp", a.zc.App.Addr)
	if err != nil {
		return err
	}

	if a.rpcServer != nil {
		if err = a.rpcServer.Start(l); err != nil {
			return err
		}
	}

	if a.httpServer != nil {
		if err = a.httpServer.Start(l); err != nil {
			return err
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	log.Infof("received signal %s", <-ch)

	if a.rpcServer != nil {
		_ = a.rpcServer.Stop()
	}

	if a.httpServer != nil {
		_ = a.httpServer.Stop()
	}

	return nil
}
