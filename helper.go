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
func AddGocronInstance(name string, f1 gocron.InitializeFunction, f2 gocron.RuntimeFunction) error {
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

// RemoveJob 移除job
func RemoveJob(name string, id string) bool {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Remove(id)
		return true
	}
	return false
}

// RemoveAllJob 移除所有job
func RemoveAllJob(name string) {
	vos := StateJob(name)
	for _, vo := range vos {
		RemoveJob(name, vo.JobId)
	}
}

// ReloadJob 重新加载job
func ReloadJob(name string, id string) {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Reload(id)
	}
}

func GetEntryItems(name string) []gocron.JobStateItem {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		return client.Prints()
	}
	return []gocron.JobStateItem{}
}

// Stop 停止gocron
func Stop(name string) {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		client.Cron().Stop()
	}
}

// StateJob 获取所有job状态
func StateJob(name string) []gocron.JobStateVO {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		return client.State()
	}
	return make([]gocron.JobStateVO, 0)
}

// GetCron 获取cron.Cron对象
func GetCron(name string) *cron.Cron {
	if HasGocronInstance(name) {
		client := GetCronClient(name)
		return client.Cron()
	}
	return nil
}
