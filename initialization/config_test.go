package initialization

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	path, _ := filepath.Abs("../config")
	viper.AddConfigPath(path)
	Config()

	fmt.Println(viper.Get("database"))
}
