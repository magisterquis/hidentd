package main

/*
 * hidentd.go
 * Yet another spoofing identd
 * By J. Stuart McMurray
 * Created 20160529
 * Last Modified 20160529
 */

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

const DEFPORT = "61113"

func main() {
	var (
		uname = flag.String(
			"u",
			"hidentd",
			"Spoofed `username`",
		)
		os = flag.String(
			"o",
			"UNIX",
			"`OS` to return",
		)
		addr = flag.String(
			"l",
			"0.0.0.0:61113",
			"Listen [address and] port",
		)
	)

	/* Make sure we have an address */
	addr = cleanAddr(*addr)

	/* Listen on the address */
	l, err := net.Listen("tcp", *addr)
	if nil != err {
		log.Fatalf("Unable to listen on %v: %v", *addr, err)
	}
	log.Printf("Listening on %v", l.Addr())

	/* Handle clients */
	for {
		c, err := l.Accept()
		if nil != err {
			log.Fatalf(
				"Unable to accept client on %v: %v",
				l.Addr(),
				err,
			)
		}
		go handle(c, *os, *uname)
	}
}

// cleanAddr makes sure the address is ready to be listened upon.
func cleanAddr(a string) *string {
	/* Make sure it's not just a number */
	if i, err := strconv.Atoi(a); nil == err {
		b := fmt.Sprintf(":%v", i)
		return &b
	}

	/* Make sure we have a port */
	_, _, err := net.SplitHostPort(a)
	if nil != err {
		if !strings.HasPrefix(err.Error(), "missing port in address") {
			log.Fatalf(
				"Unable to check for port in %q: %v",
				a,
				err,
			)
		}
		b := net.JoinHostPort(a, DEFPORT)
		return &b
	}
	return &a
}

// handle logs and replies to a client
func handle(c net.Conn, os, uname string) {
	tc := textproto.NewConn(c)

	/* Get the request */
	req, err := tc.ReadLine()
	if nil != err {
		log.Printf("Address:%v ReadError:%q", c.RemoteAddr(), err)
		return
	}

	/* Craft a reply */
	res := fmt.Sprintf(
		"%v : USERID : %v : %v",
		strings.TrimSpace(req),
		os,
		uname,
	)

	/* Send the response */
	if err := tc.PrintfLine("%v", res); nil != err {
		log.Printf("Address:%v WriteError:%v", c.RemoteAddr(), err)
	}

	log.Printf("Address:%v Request:%q Response:%q", c.RemoteAddr(), req, res)
}
