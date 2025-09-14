package bootstrap

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/flags"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initConfig() {
	_ = godotenv.Load(flags.Env)

	v := viper.New()
	v.SetConfigName("config")
	if wd, err := os.Getwd(); err == nil {
		v.AddConfigPath(wd)
	}

	_ = v.ReadInConfig()

	keys := config.GetMapStructure()
	envKeys := config.GetMapStructureForEnv(keys)
	if len(keys) == len(envKeys) {
		for i := 0; i < len(keys); i++ {
			if val, ok := os.LookupEnv(envKeys[i]); ok {
				if !v.IsSet(keys[i]) {
					v.Set(keys[i], val)
				}
			}
		}
	}

	var cfg config.Config
	_ = v.Unmarshal(&cfg)
	config.SetConfig(cfg)
}

func CreateConfigFile(configType string) error {
	// 根據類型創建對應的設定檔
	switch configType {
	case "env":
		return createEnvConfig()
	case "json":
		return createJSONConfig()
	case "yaml":
		return createYAMLConfig()
	case "toml":
		return createTOMLConfig()
	default:
		return fmt.Errorf("不支援的設定檔類型: %s，支援的類型：env, json, yaml, toml", configType)
	}
}

// 從 config.Config 結構體生成預設值
func generateDefaultConfig() config.Config {
	return config.GetDefault()
}

func createEnvConfig() error {
	cfg := generateDefaultConfig()
	keys := config.GetMapStructure()
	envKeys := config.GetMapStructureForEnv(keys)

	lines := make([]string, 0, len(keys))
	for i, key := range keys {
		value := getValueByPath(cfg, key)
		lines = append(lines, fmt.Sprintf("%s=%v", envKeys[i], value))
	}

	content := strings.Join(lines, "\n")
	return os.WriteFile(flags.Env, []byte(content), 0o644)
}

func createJSONConfig() error {
	cfg := generateDefaultConfig()

	v := viper.New()
	setConfigValues(v, cfg)

	return v.WriteConfigAs("config.json")
}

func createYAMLConfig() error {
	cfg := generateDefaultConfig()

	v := viper.New()
	setConfigValues(v, cfg)

	return v.WriteConfigAs("config.yaml")
}

func createTOMLConfig() error {
	cfg := generateDefaultConfig()

	v := viper.New()
	setConfigValues(v, cfg)

	return v.WriteConfigAs("config.toml")
}

// 使用 viper 設定值
func setConfigValues(v *viper.Viper, cfg config.Config) {
	keys := config.GetMapStructure()
	for _, key := range keys {
		value := getValueByPath(cfg, key)
		v.Set(key, value)
	}
}

// 根據路徑獲取結構體中的值
func getValueByPath(cfg config.Config, path string) interface{} {
	parts := strings.Split(path, ".")
	current := reflect.ValueOf(cfg)

	for _, part := range parts {
		if current.Kind() == reflect.Ptr {
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return nil
		}

		field := current.FieldByNameFunc(func(name string) bool {
			field, _ := current.Type().FieldByName(name)
			return field.Tag.Get("mapstructure") == part
		})

		if !field.IsValid() {
			return nil
		}

		current = field
	}

	return current.Interface()
}
