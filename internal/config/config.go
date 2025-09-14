package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/duke-git/lancet/v2/slice"
)

var (
	envReplacer = strings.NewReplacer(".", "_")
	cfg         Config
)

type Config struct {
	Log Log `mapstructure:"log"`
}

func SetConfig(c Config) {
	cfg = c
}

func GetConfig() Config {
	return cfg
}

func GetEnvReplacer() *strings.Replacer {
	return envReplacer
}

func GetMapStructureForEnv(s []string) []string {
	r := GetEnvReplacer()
	return slice.Map(s, func(_ int, s string) string {
		return strings.ToUpper(r.Replace(s))
	})
}

func GetMapStructure() []string {
	t := reflect.TypeOf(new(Config))
	return getMapStructure(t, "mapstructure")
}

func getMapStructure(t reflect.Type, tag string) []string {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if fTag := f.Tag.Get(tag); fTag != "" {
			switch f.Type.Kind() {
			case reflect.Struct:
				tags := getMapStructure(f.Type, tag)
				for _, s2 := range tags {
					fields = append(fields, fmt.Sprintf("%s.%s", fTag, s2))
				}
			default:
				fields = append(fields, fTag)
			}
		}
	}
	return fields
}

func GetDefault() Config {
	return Config{
		Log: Log{
			Level:      LogLevelInfo,
			Format:     LogFormatJSON,
			Console:    true,
			OutputPath: "./logs/log.log",
			ErrorPath:  "./logs/error.log",
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 10,
			Compress:   true,
		},
	}
}
