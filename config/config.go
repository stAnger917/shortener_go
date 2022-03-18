package config

import (
	"os"
	"shortener/pkg/logging"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP     HTTPConfig
		Provider ProviderCfg
	}

	ProviderCfg struct {
		Token  string
		URL    string
		Domain string
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}
)

func InitConfig(envPath string) *Config {
	var AppConfig Config
	AppConfig.HTTP = InitHttpConfig(envPath)
	AppConfig.Provider = InitProviderConfig(envPath)
	logging.EasyLogInfo("config", "app configuration successfully completed", "")
	return &AppConfig
}

func InitHttpConfig(envPath string) HTTPConfig {
	var AppHTTPConfig HTTPConfig

	if err := godotenv.Load(envPath); err != nil {
		logging.EasyLogError("config", "failed to init HTTP configuration! ", "building config from default values", err)
		return SetDefaultHTTPConfig()
	}

	srvHost, exists := os.LookupEnv("HTTP_HOST")
	if exists {
		AppHTTPConfig.Host = srvHost
	}

	srvPort, exists := os.LookupEnv("HTTP_PORT")
	if exists {
		AppHTTPConfig.Port = srvPort
	}

	srvRTimeout, exists := os.LookupEnv("HTTP_READ_TIMEOUT")
	if exists {
		timeOut, _ := time.ParseDuration(srvRTimeout)
		AppHTTPConfig.ReadTimeout = timeOut
	}

	srvWTimeout, exists := os.LookupEnv("HTTP_WRITE_TIMEOUT")
	if exists {
		timeOut, _ := time.ParseDuration(srvWTimeout)
		AppHTTPConfig.WriteTimeout = timeOut
	}

	srvMaxHeaderBytes, exists := os.LookupEnv("MAX_HEADER_BYTES")
	if exists {
		maxHeaderMegabytes, _ := strconv.Atoi(srvMaxHeaderBytes)
		AppHTTPConfig.MaxHeaderMegabytes = maxHeaderMegabytes
	}
	logging.EasyLogInfo("config", "initiated HTTP configuration from env", "")
	return AppHTTPConfig
}

func InitProviderConfig(envPath string) ProviderCfg {
	var Configuration ProviderCfg

	if err := godotenv.Load(envPath); err != nil {
		logging.EasyLogFatal("config", "failed to init provider configuration! ", "", err)
		return ProviderCfg{}
	}
	providerToken, exists := os.LookupEnv("BITLY_TOKEN")
	if exists {
		Configuration.Token = providerToken
	}

	providerURL, exists := os.LookupEnv("BITLY_URL")
	if exists {
		Configuration.URL = providerURL
	}
	providerDomain, exists := os.LookupEnv("BITLY_DOMAIN")
	if exists {
		Configuration.Domain = providerDomain
	}
	return Configuration
}
func SetDefaultHTTPConfig() HTTPConfig {
	var defaultTimeOutDuration time.Duration = 10
	var DefaultHTTPConfig HTTPConfig
	DefaultHTTPConfig.Host = "http://127.0.0.1"
	DefaultHTTPConfig.Port = "8000"
	DefaultHTTPConfig.ReadTimeout = defaultTimeOutDuration
	DefaultHTTPConfig.WriteTimeout = defaultTimeOutDuration
	DefaultHTTPConfig.MaxHeaderMegabytes = 1

	return DefaultHTTPConfig
}
