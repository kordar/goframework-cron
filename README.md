# goframework-cron

| 名称                                  | 方法                   | 描述                                        |
|-------------------------------------|----------------------|-------------------------------------------|
| 添加并创建`Gocron`实例                     | AddGocronInstance    | 
| 移除`Gocron`实例                        | RemoveGocronInstance | 1、移除所有job <br /> 2、停止cron线程 <br /> 3、移除引用 |
| `Gocron`实例是否存在                      | HasGocronInstance    |                                           |
| `Gocron`实例添加`Job`                   | AddJob               | 向`Gocron`对象中添加`gocron.Schedule`实例         |
| `Gocron`实例移除`Job`                   | RemoveJob            |                                           |
| `Gocron`实例移除所有`Job`                 | RemoveAllJob         |                                           |
| `Gocron`实例重新加载`Job`                 | ReloadJob         |                                           |
| `Gocron`实例`JobStateItem`列表获取          | GetEntryItems        |                                           |
| `Gocron`所有`Job`状态                   | StateJob                 |                                           |
| `Gocron`实例停止                        | Stop                 |                                           |
| 原始`*cron.Cron`对象获取                  | GetCron                 |                                           |


## 功能接口

```go
func AddGocronInstance(name string, f1 gocron.InitializeFunction, f2 gocron.RuntimeFunction) error 
func RemoveGocronInstance(name string)
func HasGocronInstance(name string) bool 
func AddJob(name string, job gocron.Schedule) bool
func RemoveJob(name string, id string) bool
func RemoveAllJob(name string)
func ReloadJob(name string, id string)
func GetEntryItems(name string) []gocron.JobStateItem 
func Stop(name string)
func StateJob(name string) []gocron.JobStateVO
func GetCron(name string) *cron.Cron
```

- 实现`Job`

```go
type TestNameSchedule struct {
    gocron.BaseSchedule
}

func (s *TestNameSchedule) GetId() string {
    return "test-name"
}

func (s *TestNameSchedule) Execute() {
    config := s.Config()
    logger.Infof("----------------------------%v", config)
}
```

- 添加`Job`  

```go
// 系统默认封装cron.FuncJob函数进行添加
AddJob("xxx", &TestNameSchedule{})
```
