package goframework_cron

import (
	"github.com/kordar/gocron"
)

type GocronIns struct {
	name string
	ins  *gocron.Gocron
}

func NewGocronIns(name string, f1 func(job gocron.Schedule) map[string]string, f2 func(job gocron.Schedule) bool) *GocronIns {
	G := gocron.NewGocron(f1, f2)
	G.Start()
	return &GocronIns{name: name, ins: G}
}

func (c GocronIns) GetName() string {
	return c.name
}

func (c GocronIns) GetInstance() interface{} {
	return c.ins
}

func (c GocronIns) Close() error {
	return nil
}
