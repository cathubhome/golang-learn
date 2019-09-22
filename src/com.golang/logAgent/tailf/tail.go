package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"sync"
	"time"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

func InitTail(conf []CollectConf, chanSize int) (err error) {
	//实例化tailObjMgr优先于len(conf)代码段,否则com.golang/logAgent/tailf.GetMsgLine(...)报错
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(conf) == 0 {
		logs.Warn("cannot find any etcd log collect configuration")
		return
	}

	for _, v := range conf {
		createNewTask(v)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailObj.conf.Topic,
			}

			tailObjMgr.msgChan <- textMsg
			//不再采集的日志
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf:%v", tailObj.conf)
			return
		}
	}
}

func GetMsgLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

//更新配置，这里在多个goroutine里面操作tailObjMgr.tailObj，要加互斥锁
func UpdateConfig(confs []CollectConf) {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()

	//刚开始进行etcd的日志配置，或对收集日志的配置发生新增时
	for _, confItem := range confs {
		var isRunning = false
		//遍历正在运行的tailf实例
		for _, obj := range tailObjMgr.tailObjs {
			//日志的全路径比对，相同说明日志已经在收集了
			if confItem.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}
		if isRunning {
			continue
		}
		createNewTask(confItem)
	}

	//删除了etcd的日志配置（键值对），或对收集日志的配置发生删减时
	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		//通过channel通信,对不再需要收集的日志实例进行退出
		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		//存留下来的需要收集的日志配置实例
		tailObjs = append(tailObjs, obj)
	}

	tailObjMgr.tailObjs = tailObjs
	return
}

//收集日志任务
func createNewTask(conf CollectConf) {
	fmt.Printf("create new task for collecting logs:%v\n", conf)
	logs.Info("create new task for CollectConf:%v\n", conf)
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})

	if errTail != nil {
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, errTail)
		return
	}

	obj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFromTail(obj)

}

/**
收集日志配置的结构体
*/
type CollectConf struct {
	LogPath string `json:"path"`
	Topic   string `json:"topic"`
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf

	status   int
	exitChan chan int
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

type TextMsg struct {
	Msg   string
	Topic string
}

var (
	tailObjMgr *TailObjMgr
)
