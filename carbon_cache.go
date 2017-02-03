// Listen to connections, receiving metrics

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
	metricCache		*MemCache

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

}


func main() {
	
}