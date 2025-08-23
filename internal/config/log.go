package config

type LogLevel = string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type LogFormat = string

const (
	LogFormatJson    LogFormat = "json"
	LogFormatConsole LogFormat = "console"
)

type Log struct {
	Level      LogLevel  `mapstructure:"level"`
	Format     LogFormat `mapstructure:"format"`
	Console    bool      `mapstructure:"console"`
	OutputPath string    `mapstructure:"output_path"`
	ErrorPath  string    `mapstructure:"error_path"`
	MaxSize    int       `mapstructure:"max_size"`
	MaxAge     int       `mapstructure:"max_age"`
	MaxBackups int       `mapstructure:"max_backups"`
	Compress   bool      `mapstructure:"compress"`
}
