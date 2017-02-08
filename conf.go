package main

// 1. Config Struct
// 2. Instance Struct
// 3. Load config

type Config struct {
	STORAGE_DIR             string
	DATABASE                string
	MAX_CACHE_SIZE          float64
	workers                 map[string][]*Instance
	CACHE_WRITE_STRATEGY    string
}


type Instance struct {
	name                    string
	// PICKLE_RECEIVER_PORT    int
	LINE_RECEIVER_PORT	    int
	CACHE_QUERY_PORT        int
}



// load config file and construct a Config object