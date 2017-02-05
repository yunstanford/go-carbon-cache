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
)


var (
	metricCache    *MemCache

)

func accept(l *net.TCPListener, config Config) {
	for {
		c, err := l.AcceptTCP()
		if nil != err {
			// log.Error(err.Error())
			break
		}
		go handle(c, config)
	}
}


func handle(c net.Conn, config Config) {
	defer c.Close()
	// receive metrics and store them in MemCache


}


func drain() {

}

func start(config Config) {
	

	// Starting Writer
	go drain()
	// Starting Listener
	
}