package config

type K8s struct {
	Namespace       string `mapstructure:"namespace"`
	URI             string `mapstructure:"uri"`
	ImagePullSecret string `mapstructer:"image_pull_secret"`
}
