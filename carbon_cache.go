// Listen to connections, receiving metrics
// Protocol for one carbon cache instance

package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"syscall"
	"time"
	"bufio"
	"github.com/rcrowley/goagain"
	logging "github.com/op/go-logging"
	"github.com/BurntSushi/toml"

)


var log = logging.MustGetLogger("carbon-relay-ng")

var (
	metricCache    *MemCache

)

func accept(l *net.TCPListener, config Config) {
	for {
		c, err := l.AcceptTCP()
		if nil != err {
			log.Error(err.Error())
			break
		}
		go handle(c, config)
	}
}


func handle(c net.Conn, config Config) {
	defer c.Close()
	// receive metrics and store them in MemCache,
	// Process as fast as possible.
	// Drop received metrics if memcache is full.
	reader := bufio.NewReaderSize(c, 4096)   // let's pick up the same value as carbon-relay-ng
	for {
		buf, _, err := r.ReadLine()

		if nil != err {
			if io.EOF != err {
				log.Error(err.Error())
			}
			break
		}

		// TODO: check if cache is full

		// TODO: Store metric in memcache
		
	}

}


func drain() {
	for {

	}
}


func start(config tomlConfig, worker string) {
	
	
	// Starting Writer
	go drain()
	// Starting Listener
	
}