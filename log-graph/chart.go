package loggraph

type ChartConfig struct {
	Name  string `yaml:"name"`
	Regex string `yaml:"regex"`
	XAxis string `yaml:"x_axis"`
	YAxis string `yaml:"y_axis"`
	Color string `yaml:"color"`
}

type Config struct {
	Port   int           `yaml:"port"`
	Charts []ChartConfig `yaml:"charts"`
}
