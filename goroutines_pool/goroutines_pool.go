package main

import (
	"fmt"
	"sync"
	"time"
)

//go协程池
type GoroutinesPool struct {
	WorkerFunc     func(int) error //工作函数
	MaxWorkerCount int             //最大连接数
	workdersCount  int
	lock           sync.Mutex
	ready          chan *WorkerChan
	accpetChan     chan int
	workerChanPool sync.Pool //chan poolm  tft
	waitWorker     bool
}

//工作Chan
type WorkerChan struct {
	lastUseTime time.Time //上次使用时间
	ch          chan int
}

func NewGoroutinesPool(f func(int) error, MaxWorkerCount int) *GoroutinesPool {
	gp := &GoroutinesPool{
		WorkerFunc:     f,
		MaxWorkerCount: MaxWorkerCount,
		accpetChan:     make(chan int),
	}
	gp.ready = make(chan *WorkerChan, MaxWorkerCount)
	go gp.Loop()
	return gp
}

func (gp *GoroutinesPool) GoWork(cmd int) {
	gp.accpetChan <- cmd
}

func (gp *GoroutinesPool) server(cmd int) bool {
	ch := gp.getCh()
	if ch == nil {
		return false
	}
	ch.ch <- cmd
	return true
}

//获取一个chan
func (gp *GoroutinesPool) getCh() *WorkerChan {
	var ch *WorkerChan
	waitChan := false
	gp.lock.Lock()

	ready := gp.ready
	n := len(ready) - 1
	if n < 0 {
		if gp.workdersCount < gp.MaxWorkerCount {
			gp.workdersCount++
		} else {
			waitChan = true
		}
	} else {
		ch = <-gp.ready
		// ch = ready[n]
		// ready[n] = nil
		// gp.ready = ready[:n]
		// gp.waitWorker = false
	}
	gp.lock.Unlock()
	if waitChan {
		ch = <-gp.ready
	}
	if ch == nil {
		// if !createWorker {
		// 	return nil
		// }
		//创建新的协程
		vch := gp.workerChanPool.Get()
		if vch == nil {
			vch = &WorkerChan{
				ch: make(chan int, 1),
			}
		}
		ch = vch.(*WorkerChan)
		go func() {
			gp.workerFunc(ch)
			gp.workerChanPool.Put(vch)
		}()
	}
	return ch
}

func (gp *GoroutinesPool) workerFunc(ch *WorkerChan) {
	for cmd := range ch.ch {
		//释放goroutines
		if cmd == -1 {
			break
		}
		//真正处理的函数
		gp.WorkerFunc(cmd)
		//释放chan
		gp.release(ch)
	}
	gp.lock.Lock()
	gp.workdersCount--
	gp.lock.Unlock()
}

//释放workerChan
func (gp *GoroutinesPool) release(ch *WorkerChan) bool {
	ch.lastUseTime = time.Now()
	gp.lock.Lock()
	defer gp.lock.Unlock()
	gp.ready <- ch
	//gp.ready = append(gp.ready, ch)
	return true
}

//循环
func (gp *GoroutinesPool) Loop() {

	for {
		cmd := <-gp.accpetChan
		//处理
		if !gp.server(cmd) {
			fmt.Println("max goroutinePool")
		}
	}
}
