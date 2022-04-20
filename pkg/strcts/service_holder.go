package strcts

import "time"

type AppConfig struct {
	Clusters         []Cluster     `yaml:"clusters"`
	SlackUrl         string        `yaml:"slackUrl"`
	EvaluateInterval time.Duration `yaml:"evaluateInterval"`
}
