package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/VictoriaMetrics/metrics"
	"github.com/aelsabbahy/GOnetstat"
	"github.com/coreos/go-systemd/daemon"
)

const (
	appName         = "sockstat_exporter"
	metricName      = "sockstat_listen"
	metricErrorName = "sockstat_errors"
)

var (
	revision   string
	appVersion = appName + " " + revision

	listen string
	ver    bool
)

func init() {
	flag.StringVar(&listen, "listen", ":9997", "Listen metrics server address. [env: LISTEN]")
	flag.BoolVar(&ver, "v", false, "Print version")
	flag.Parse()
	if ver {
		fmt.Println(appVersion)
		os.Exit(0)
	}
	envListen := os.Getenv("LISTEN")
	if envListen != "" {
		listen = envListen
	}
}

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		s := metrics.NewSet()
		wg := sync.WaitGroup{}
		wg.Add(1)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}
			tcp, err := GOnetstat.Tcp(true)
			if err != nil {
				fmt.Println(err)
				s.GetOrCreateCounter(metricErrorName).Inc()
			}
			tcp6, err := GOnetstat.Tcp6(true)
			if err != nil {
				fmt.Println(err)
				s.GetOrCreateCounter(metricErrorName).Inc()
			}
			udp, err := GOnetstat.Udp(true)
			if err != nil {
				fmt.Println(err)
				s.GetOrCreateCounter(metricErrorName).Inc()
			}
			udp6, err := GOnetstat.Udp6(true)
			if err != nil {
				fmt.Println(err)
				s.GetOrCreateCounter(metricErrorName).Inc()
			}
			for _, p := range tcp {
				if p.State == "LISTEN" {
					name := fmt.Sprintf(`%s{proto="tcp",address="%s",port="%d",path="%s"}`, metricName, p.Ip, p.Port, p.Exe)
					s.GetOrCreateGauge(name, func() float64 { return 1 })
				}
			}
			for _, p := range tcp6 {
				if p.State == "LISTEN" {
					name := fmt.Sprintf(`%s{proto="tcp6",address="%s",port="%d",path="%s"}`, metricName, p.Ip, p.Port, p.Exe)
					s.GetOrCreateGauge(name, func() float64 { return 1 })
				}
			}
			for _, p := range udp {
				if p.State == "CLOSE" {
					name := fmt.Sprintf(`%s{proto="udp",address="%s",port="%d",path="%s"}`, metricName, p.Ip, p.Port, p.Exe)
					s.GetOrCreateGauge(name, func() float64 { return 1 })
				}
			}
			for _, p := range udp6 {
				if p.State == "CLOSE" {
					name := fmt.Sprintf(`%s{proto="udp6",address="%s",port="%d",path="%s"}`, metricName, p.Ip, p.Port, p.Exe)
					s.GetOrCreateGauge(name, func() float64 { return 1 })
				}
			}
		}()
		wg.Wait()
		s.WritePrometheus(w)
	})
	daemon.SdNotify(false, "READY=1")
	log.Fatal(http.ListenAndServe(listen, nil))
}
