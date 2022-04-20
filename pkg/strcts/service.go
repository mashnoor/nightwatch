package strcts

type Cluster struct {
	ClusterName       string  `yaml:"clusterName"`
	DbUser            string  `yaml:"dbUser"`
	DbHost            string  `yaml:"dbHost"`
	DbPort            int     `yaml:"dbPort"`
	DbPassword        string  `yaml:"dbPassword"`
	LogDelayThreshold float32 `yaml:"logDelayThreshold"`
}
