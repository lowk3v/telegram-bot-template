package configs

import (
	"crypto/tls"
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type secretConfig struct {
	EnvMode string `json:"PROJECT_NAME_ENV"`
}

type appConfig struct {
}

type SymbolConfig struct {
	Success string
	Error   string
	Info    string
}

var Secret secretConfig
var AppConfig appConfig
var HttpClient *http.Client
var Symbol SymbolConfig

//go:embed config.yaml
var appConfigYaml string

//go:embed *.env
var envFile string

func init() {
	// load environment from file. Be overridden by system environment
	envContent, err := godotenv.Unmarshal(envFile)
	err = convertToStruct(envContent, &Secret)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", color.RedString("¿"), err)
	}
	loadSystemEnv(&Secret)

	// load config based on environment (production, development, testnet)
	var appConfigTemp map[string]interface{}
	err = yaml.Unmarshal([]byte(appConfigYaml), &appConfigTemp)
	err = convertToStruct(appConfigTemp[Secret.EnvMode], &AppConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", color.RedString("¿"), err)
	}

	Symbol = SymbolConfig{
		Success: color.GreenString("≠"),
		Error:   color.RedString("¿"),
		Info:    color.BlueString("ℹ"),
	}
}

func loadSystemEnv(s *secretConfig) {
	if len(os.Getenv("PROJECT_NAME_ENV")) != 0 {
		s.EnvMode = os.Getenv("SCON_ENV")
	}
	// insert more if handle more environment variables
}

func convertToStruct(from interface{}, to interface{}) error {
	jsonS, err := json.Marshal(from)
	err = json.Unmarshal(jsonS, to)
	if err != nil {
		return err
	}
	return nil
}

func InitHttpClient(proxy string) {
	transport := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	if proxy != "" {
		if p, err := url.Parse(proxy); err == nil {
			transport.Proxy = http.ProxyURL(p)
		}
	}

	redirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	HttpClient = &http.Client{
		Transport:     transport,
		CheckRedirect: redirect,
		Timeout:       time.Second * 10,
	}
}

func (s *secretConfig) PrintInfo() {
	fmt.Printf("CURRENT ENVIRONMENT: [%s]. To switch environment, set PROJECT_NAME_ENV=...\n\n",
		color.BlueString(Secret.EnvMode))
}
