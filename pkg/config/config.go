package config

import (
    "fmt"
    "os"
    
    "gopkg.in/yaml.v2"
)

type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Dingtalk DingtalkConfig `yaml:"dingtalk"`
    Database DatabaseConfig `yaml:"database"`
    Redis    RedisConfig    `yaml:"redis"`
}

type ServerConfig struct {
    Address string `yaml:"address"`
    Mode    string `yaml:"mode"`
}

type DingtalkConfig struct {
    AppKey    string `yaml:"appKey"`
    AppSecret string `yaml:"appSecret"`
    AgentId   string `yaml:"agentId"`
}

type DatabaseConfig struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
    Addr     string `yaml:"addr"`
    Password string `yaml:"password"`
    DB       int    `yaml:"db"`
}

var config Config

func Get() *Config {
    return &config
}

func Load(env string) error {
    configPath := fmt.Sprintf("configs/config.%s.yaml", env)
    data, err := os.ReadFile(configPath)
    if err != nil {
        return fmt.Errorf("读取配置文件失败: %v", err)
    }

    if err := yaml.Unmarshal(data, &config); err != nil {
        return fmt.Errorf("解析配置文件失败: %v", err)
    }

    return nil
} 