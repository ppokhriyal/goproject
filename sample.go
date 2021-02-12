package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	vi := viper.New()
	vi.SetConfigFile("sample.yml")
	vi.ReadInConfig()
	fmt.Println(vi.GetString("awsvpc.name"))
}