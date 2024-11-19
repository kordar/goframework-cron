package goframework_rabbitmq

import (
	"github.com/kordar/gocron"
	"github.com/kordar/godb"
	"github.com/robfig/cron/v3"
)

var (
	cronpool = godb.NewDbPool()
)

func GetCronClient(db string) *gocron.Gocron {
	return cronpool.Handle(db).(*gocron.Gocron)
}

// AddGocronInstance 添加cron
func AddGocronInstance(db string, f1 func(job gocron.Schedule) map[string]string, f2 func(job gocron.Schedule) bool) error {
	ins := NewGocronIns(db, f1, f2)
	return cronpool.Add(ins)
}

// RemoveGocronInstance 移除cron
func RemoveGocronInstance(db string) {
	cronpool.Remove(db)
}

// HasGocronInstance cron句柄是否存在
func HasGocronInstance(db string) bool {
	return cronpool != nil && cronpool.Has(db)
}

func Add(db string, job gocron.Schedule) bool {
	if HasGocronInstance(db) {
		client := GetCronClient(db)
		client.Remove(job.GetId())
		client.Add(job)
		return true
	}
	return false
}

func AddWithJob(db string, job gocron.Schedule, funcJob cron.Job) bool {
	if HasGocronInstance(db) {
		client := GetCronClient(db)
		client.Remove(job.GetId())
		client.AddWithJob(job, funcJob)
		return true
	}
	return false
}

func AddWithChain(db string, job gocron.Schedule, f func(funcJob cron.Job) cron.Job) bool {
	if HasGocronInstance(db) {
		client := GetCronClient(db)
		client.Remove(job.GetId())
		client.AddWithChain(job, f)
		return true
	}
	return false
}

func Stop(db string, id string) bool {
	if HasGocronInstance(db) {
		client := GetCronClient(db)
		client.Remove(id)
		return true
	}
	return false
}

func Prints(db string) []*gocron.EntryItem {
	if HasGocronInstance(db) {
		client := GetCronClient(db)
		return client.Prints()
	}
	return []*gocron.EntryItem{}
}
