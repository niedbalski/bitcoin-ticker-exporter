package main

import
(
	"github.com/niedbalski/bitcoin-ticker-exporter/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/op/go-logging.v1"
	"strings"
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/niedbalski/bitcoin-ticker-exporter/exporters"
)

type Logger interface {
	Debugf(format string, args ...interface{})
}


var Log Logger = logging.MustGetLogger("bitcoin-ticker-exporter")


type stringList []string

func (i *stringList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *stringList) String() string {
	return strings.Join(*i, ", ")
}

func (i *stringList) IsCumulative() bool {
	return true
}

func StringList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*stringList)(target))
	return
}

func main()() {
	var (
		bind = kingpin.Flag("web.listen-address", "address:port to listen on").Default(":9180").String()
		metrics = kingpin.Flag("web.telemetry-path", "uri path to expose metrics").Default("/metrics").String()
		exchanges = StringList(kingpin.Flag("exchange", "enable ticker exporting for given exchange").Required())
		configFile = kingpin.Flag("config", "Path to configuration file").Default("./settings.yml").String()
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	settings, err := config.NewConfigFromYAML(*configFile)
	if err != nil {
		fmt.Println(err)
	}

	exporter, err := exporters.NewExporter(settings)

	if err != nil {
		fmt.Println("Error", exchanges, bind)
	}


	prometheus.MustRegister(exporter)


	http.Handle(*metrics, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Memcached Exporter</title></head>
             <body>
             <h1>Bitcoin Exporter</h1>
             <p><a href='` + *metrics + `'>Metrics</a></p>
             </body>
             </html>`))
	})


	log.Infoln("Starting HTTP server on", *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))

}
