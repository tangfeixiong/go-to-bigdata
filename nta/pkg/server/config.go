package server

import (
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/hbase"
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/config"
)

type Config struct {
	Common *config.Config

	HBase hbase.Config
}
