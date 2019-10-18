package main

import (
	"github.com/yinghuocho/gosocks"
	"time"
)

func main() {
	srv := gosocks.NewBasicServer("127.0.0.1:20800", time.Minute)
	srv.ListenAndServe()
}
