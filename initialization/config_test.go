package initialization

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	path, _ := filepath.Abs("../config")
	viper.AddConfigPath(path)
	Config()

	fmt.Println(viper.Get("database"))
}
