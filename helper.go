package goframework_cron

import (
	"github.com/kordar/gocron"
	"github.com/kordar/godb"
	"github.com/robfig/cron/v3"
)

var (
	cronpool = godb.NewDbPool()
)

func GetCronClient(name string) *gocron.Gocron {
	return cronpool.Handle(name).(*gocron.Gocron)
}

// AddGocronInstance 添加cron
func AddGocronInstance(name string, f1 func(job gocron.Schedule) map[string]string, f2 func(job gocron.Schedule) bool) error {
	ins := NewGocronIns(name, f1, f2)
	return cronpool.Add(ins)
}

// RemoveGocronInstance 移除cron
func RemoveGocronInstance(name string) {
	RemoveAllJob(name)
	Stop(name)
	cronpool.Remove(name)
}

// HasGocronInstance cron句柄是否存在
func HasGocronInstance(name string) bool {
	return cronpool != nil && cronpool.Has(name)
}

func AddJob(name string, job gocron.Schedule) bool {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Remove(job.GetId())
		client.Add(job)
		return true
	}
	return false
}

func AddCronJob(name string, job gocron.Schedule, funcJob cron.Job) bool {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Remove(job.GetId())
		client.AddWithJob(job, funcJob)
		return true
	}
	return false
}

func AddCronJobWithChain(name string, job gocron.Schedule, f func(funcJob cron.Job) cron.Job) bool {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Remove(job.GetId())
		client.AddWithChain(job, f)
		return true
	}
	return false
}

func RemoveJob(name string, id string) bool {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Remove(id)
		return true
	}
	return false
}

func RemoveAllJob(name string) {
	items := GetEntryItems(name)
	for _, item := range items {
		RemoveJob(name, item.Id)
	}
}

func GetEntryItems(name string) []*gocron.EntryItem {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		return client.Prints()
	}
	return []*gocron.EntryItem{}
}

func Stop(name string) {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Stop()
	}
}

// GetCron 获取cron.Cron对象
func GetCron(name string) *cron.Cron {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		return client.Cron()
	}
	return nil
}
