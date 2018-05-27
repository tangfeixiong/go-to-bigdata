package server

import (
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/hbase"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	HBase           hbase.Config
}
