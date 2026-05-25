package main

import (
    "fmt"
    "os"
)

type Config struct {
    // AWS CONFIG
    AWSRegion string
    AWSKey    string
    AWSSecret string
    AWSBucket string

    // DATABASE CONFIG
    DBHost     string
    DBPort     string
    DBUser     string
    DBPass     string
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)

    if value == "" {
        if defaultValue == "" {
            panic(fmt.Sprintf("Environment variable %s is required but not set", key))
        }
        return defaultValue
    }
    return value
}

func LoadConfig() Config {
    return Config {
        // aws
        AWSRegion: getEnv("AWS_REGION", "us-east-1"),
        AWSKey: getEnv("AWS_ACCESS_KEY_ID", ""),
        AWSSecret: getEnv("AWS_BUCKET", ""),
        AWSBucket: getEnv("AWS_BUCKET", ""),
        // database
        DBHost: getEnv("DB_HOST", "host.docker.internal"),
        DBPort: getEnv("DB_PORT", "3306"),
        DBUser: getEnv("DB_USER", ""),
        DBPass: getEnv("DB_PASS", ""),
    }
}
