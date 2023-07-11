package configs

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/author_name/project_name/internal/constant/const_app"
	"github.com/author_name/project_name/pkg/log"
	"github.com/author_name/project_name/pkg/pubsub/publisher"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type secretConfig struct {
	Env   string `json:"PROJECT_NAME_ENV"`
	Token string `json:"PROJECT_NAME_TOKEN"`
}

type appConfig struct {
}

type UserInteractionConfig struct {
	Sync             *sync.Mutex
	Timeout          int
	WaitReply        bool
	WaitReplyCommand string
	MetaData         map[string]interface{}
}

var Secret secretConfig
var AppConfig appConfig

//go:embed config.yaml
var appConfigYaml string

//go:embed .env
var envFile string

var GlobalPublisher *publisher.Publisher
var UserInteraction *UserInteractionConfig
var Log *log.Logger

func init() {
	Log = log.New("info")

	// load environment from file. Be overridden by system environment
	envContent, err := godotenv.Unmarshal(envFile)
	err = convertToStruct(envContent, &Secret)
	if err != nil {
		Log.WithField("env", Secret.Env).Error("Error loading environment")
	}
	loadSystemEnv(&AppConfig, &Secret)

	// load config based on environment (production, development, testnet)
	var appConfigTemp map[string]interface{}
	err = yaml.Unmarshal([]byte(appConfigYaml), &appConfigTemp)
	err = convertToStruct(appConfigTemp[Secret.Env], &AppConfig)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", color.RedString("¿"), err)
		Log.WithField("app config", Secret.Env).Error("Error loading app config")
	}

	if Secret.Env == const_app.DEVELOP {
		Log = log.New("debug")
	}

	// Queue
	GlobalPublisher = publisher.New()
	GlobalPublisher.Run()

	// User interaction
	UserInteraction = &UserInteractionConfig{
		Sync:             &sync.Mutex{},
		Timeout:          30,
		WaitReply:        false,
		WaitReplyCommand: "",
		MetaData:         make(map[string]interface{}),
	}
}

func loadSystemEnv(a *appConfig, s *secretConfig) {
	if len(os.Getenv("PROJECT_NAME_ENV")) != 0 {
		s.EnvMode = os.Getenv("PROJECT_NAME_ENV")
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

func (s *secretConfig) PrintInfo() {
	Log.WithField("environment", color.BlueString(Secret.Env)).
		Info("App Config")
}
