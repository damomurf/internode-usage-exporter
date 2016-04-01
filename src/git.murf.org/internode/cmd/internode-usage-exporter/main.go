package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/yaml.v2" // Exports as "yaml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Configuration

type Config struct {
	Username string
	Password string
}

// Internode Usage API XML Structure

type Resource struct {
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
	Name string `xml:",chardata"`
}

type Resources struct {
	Count     uint       `xml:"count,attr"`
	Resources []Resource `xml:"resource"`
}

type Service struct {
	Type    string `xml:"type,attr"`
	Request string `xml:"request,attr"`
	Href    string `xml:"href,attr"`
	Id      string `xml:",chardata"`
}

type Services struct {
	Count    uint      `xml:"count,attr"`
	Services []Service `xml:"service"`
}

type Traffic struct {
	Name         string `xml:"name,attr"`
	Rollover     string `xml:"rollover,attr"`
	PlanInterval string `xml:"plan-interval,attr"`
	Quota        uint64 `xml:"quota,attr"`
	Unit         string `xml:"unit,attr"`
	Used         uint64 `xml:",chardata"`
}

type Api struct {
	Services  Services  `xml:"services"`
	Service   Service   `xml:"service"`
	Traffic   Traffic   `xml:"traffic"`
	Resources Resources `xml:"resources"`
}

type Internode struct {
	Api Api `xml:"api"`
}

var (
	quota = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "internode",
			Subsystem: "usage",
			Name:      "quota",
			Help:      "The total quota for this service in bytes.",
		},
		[]string{
			"type",
			"id",
		},
	)
	used = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "internode",
			Subsystem: "usage",
			Name:      "used",
			Help:      "The data used for this service in bytes.",
		},
		[]string{
			"type",
			"id",
		},
	)
	target = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "internode",
			Subsystem: "usage",
			Name:      "target",
			Help:      "The target usage figure in bytes.",
		},
		[]string{
			"type",
			"id",
		},
	)
)

var (
	configPath = flag.String("config", "config.yaml", "Path to YAML configuration file")
	config     *Config
)

const API_BASE = "https://customer-webtools-api.internode.on.net/api/v1.5"

func parseConfiguration(configFile string) (*Config, error) {

	log.Printf("Looking for config in file: %s", configFile)

	f, err := os.Open(*configPath)
	if err != nil {
		return nil, err
	}

	configData, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}
	log.Printf("Read config: %+v", config)
	log.Printf("Using Internode username: %s", config.Username)

	return config, nil
}

func generateRequest(url string) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Username, config.Password)
	return req, nil
}

func discoverServices() ([]Service, error) {

	req, err := generateRequest(API_BASE)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Yes, believe it or not, 500 is what always seems to be returned here with a valid body
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		return nil, fmt.Errorf("Unexpected status code as response: %+v", resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)

	var data Internode
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	log.Printf("%+v", data)

	return data.Api.Services.Services, nil

}

func calculateTarget(quota uint64, rolloverDate string) (uint64, error) {

	layout := "2006-01-02"
	rollover, err := time.Parse(layout, rolloverDate)
	if err != nil {
		return 0, err
	}

	start := time.Date(rollover.Year(), rollover.Month()-1, rollover.Day(), 0, 0, 0, 0, time.UTC)
	quotaPerInterval := float64(quota) / float64((rollover.Sub(start)))
	target := uint64(quotaPerInterval * float64(time.Now().Sub(start)))

	return target, nil
}

func checkUsage() error {

	log.Print("Polling for usage")

	services, err := discoverServices()
	if err != nil {
		return err
	}

	for _, service := range services {

		req, err := generateRequest(fmt.Sprintf("%s/%s/usage", API_BASE, service.Href))
		if err != nil {
			return err
		}

		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Unexpected status code as response: %+v", resp.Status)
		}

		content, err := ioutil.ReadAll(resp.Body)

		var data Internode
		err = xml.Unmarshal(content, &data)
		if err != nil {
			return err
		}
		log.Printf("%+v", data)

		targetValue, err := calculateTarget(data.Api.Traffic.Quota, data.Api.Traffic.Rollover)
		if err != nil {
			return err
		}

		quota.With(prometheus.Labels{"type": service.Type, "id": service.Id}).Set(float64(data.Api.Traffic.Quota))
		used.With(prometheus.Labels{"type": service.Type, "id": service.Id}).Set(float64(data.Api.Traffic.Used))
		target.With(prometheus.Labels{"type": service.Type, "id": service.Id}).Set(float64(targetValue))

	}

	log.Print("Polling successful")
	return nil
}

func main() {

	flag.Parse()

	var err error
	config, err = parseConfiguration(*configPath)
	if err != nil {
		panic(err)
	}

	http.Handle("/metrics", prometheus.Handler())

	pollDuration := 5 * time.Minute
	log.Printf("Will poll every %s", pollDuration)

	ticker := time.NewTicker(pollDuration)
	quit := make(chan struct{})

	checkUsage()

	go func() {
		for {
			select {
			case <-ticker.C:
				err = checkUsage()
				if err != nil {
					log.Printf("Error checking usage: %+v", err)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}

	}()

	err = http.ListenAndServe(":9099", nil)
	if err != nil {
		panic(err)
	}

}

func init() {
	prometheus.MustRegister(quota)
	prometheus.MustRegister(used)
	prometheus.MustRegister(target)
}
