package main


import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// 1. Config Struct
// 2. Instance Struct
// 3. Load config

type configuration struct {
	STORAGE_DIR             string
	DATABASE                string
	MAX_CACHE_SIZE          int
	CACHE_WRITE_STRATEGY    string
}


type instance struct {
	// PICKLE_RECEIVER_PORT    int
	LINE_RECEIVER_PORT	    int
	CACHE_QUERY_PORT        int
}


// load config file and construct a Config object
// func main() {
// 	var config tomlConfig
// 	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }