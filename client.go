package guppy

import (
	"github.com/fingerpeople/guppy/config"
	"github.com/fingerpeople/guppy/libs/etcd"
)

type Guppy struct {
	Urls        []string
	DialTimeOut int
}

// Client ...
func Client(config *config.Configs) *Guppy {
	cfg := &Guppy{}
	cfg.DialTimeOut = config.DialTimeOut
	cfg.Urls = config.Urls
	return cfg
}

// New ...
func (cl *Guppy) New() (*etcd.EtcdLibs, error) {
	cfg := &config.Configs{}
	cfg.DialTimeOut = cl.DialTimeOut
	cfg.Urls = cl.Urls
	client, err := etcd.Init(cfg)
	return client, err
}
