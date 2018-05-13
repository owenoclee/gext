package config

import "log"

type Env map[string]string

func NewEnv(env map[string]string) Env {
	return Env(env)
}

func (env Env) Read(key string) string {
	if val, ok := env[key]; ok {
		return val
	}
	log.Fatalf("environment variable '%v' is required", key)
	return ""
}
