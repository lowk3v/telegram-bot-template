package configs

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/author_name/project_name/constant"
	"github.com/author_name/project_name/pkg/log"
	"github.com/author_name/project_name/pkg/pubsub/publisher"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type appConfig struct {
	Token     string `json:"token"`
	Whitelist string `json:"whitelist"`
}

type UserInteractionConfig struct {
	Sync             *sync.Mutex
	Timeout          int
	WaitReply        bool
	WaitReplyCommand string
	MetaData         map[string]interface{}
}

var AppConfig appConfig

//go:embed config.yaml
var appConfigYaml string
var GlobalPublisher *publisher.Publisher
var Log *log.Logger
var Env string

func init() {
	Log = log.New("info")

	// load config based on environment (production, development, testnet)
	var appConfigTemp map[string]interface{}
	err := yaml.Unmarshal([]byte(appConfigYaml), &appConfigTemp)
	Env = appConfigTemp["env"].(string)
	err = convertToStruct(appConfigTemp[Env], &AppConfig)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", constant.SYMBOL_ERROR, err)
		Log.WithField("app config", Env).Error("Error loading app config")
	}

	if Env == constant.DEVELOP {
		Log = log.New("debug")
	}

	// Queue
	GlobalPublisher = publisher.New()
	GlobalPublisher.Run()
}

func convertToStruct(from interface{}, to interface{}) error {
	jsonS, err := json.Marshal(from)
	err = json.Unmarshal(jsonS, to)
	if err != nil {
		return err
	}
	return nil
}

func (s *appConfig) PrintInfo() {
	Log.WithField("environment", color.BlueString(Env)).
		Info("App Config")
}
