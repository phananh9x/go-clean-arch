package config

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

//AppConfig ...
type AppConfig struct {
	Authen      AuthenConfig `yaml:"auth"`
	Mgo         *MongoConfig `yaml:"mgo"`
	Redis       *RedisConfig `yaml:"redis"`
	Kafka       *KafkaConfig `yaml:"kafka"`
	Environment string       `yaml:"environment"`
	ServiceName string       `yaml:"service_name"`
	Swagger     *Swagger     `yaml:"swagger"`
}

// AuthenConfig ...
type AuthenConfig struct {
	APIAuthenticator   AuthenticatorConfig `yaml:"api"`
	OAuth2ClientID     string              `yaml:"oauth2_client_id"`
	OAuth2ClientSecret string              `yaml:"oauth2_client_secret"`
}

// AuthenticatorConfig ...
type AuthenticatorConfig struct {
	JWTPublicKeyBase64  string `yaml:"jwt_public_key"`
	JWTPrivateKeyBase64 string `yaml:"jwt_private_key"`
	JWTExpireTime       int64  `yaml:"jwt_expire_time"`
}

// MongoConfig struct
type MongoConfig struct {
	MgoUri      string `yaml:"mgo_uri"`
	MaxPoolSize uint64 `yaml:"max_pool" default:"500"`
}

// RedisConfig ...
type RedisConfig struct {
	CacheTime           int    `yaml:"cache_time"`
	ConnectionURL       string `yaml:"connection_url"`
	PoolSize            int    `yaml:"pool_size"`
	DialTimeoutSeconds  int    `yaml:"dial_timeout_seconds"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds"`
	IdleTimeoutSeconds  int    `yaml:"idle_timeout_seconds"`
}

//KafkaConfig ...
type KafkaConfig struct {
	Addrs                     []string `yaml:"addrs"`
	Topics                    []string `yaml:"topics"`
	Group                     string   `yaml:"group"`
	MaxMessageBytes           int      `yaml:"max_message_bytes"`
	Compress                  bool     `yaml:"compress"`
	Newest                    bool     `yaml:"newest"`
	Version                   string   `yaml:"version"`
	ConsumerSessionTimout     int      `yaml:"consumer_session_timout"`
	ConsumerHeartbeatInterval int      `yaml:"consumer_heartbeat_interval"`
	ConsumerMaxProcessingTime int      `yaml:"consumer_max_processing_time"`
}

type Swagger struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"base_path"`
}

// Load load config from file and environment variables.
func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}

	fields := []interface{}{
		"func",
		"config.readFromFile",
		"filePath",
		filePath,
	}

	sugar := zap.S().With(fields...)

	sugar.Debug("Load config...")
	zap.S().Debugf("CONFIG_FILE=%v", filePath)

	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		sugar.Error("Failed to load config file")
		return nil, err
	}
	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &AppConfig{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		sugar.Error("Failed to parse config file")
		return nil, err
	}

	zap.S().Debugf("config: %+v", cfg)
	zap.S().Debug("======================================")
	zap.S().Debugf("database config: %+v", cfg.Mgo)
	zap.S().Debugf("redis config: %+v", cfg.Redis)

	return cfg, nil
}
