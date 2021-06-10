package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
)




type Worker struct {
	Client *clientv3.Client
	Ctx context.Context
	WG *sync.WaitGroup
}

func NewWorker(Config *config, ctx context.Context, wg *sync.WaitGroup) (worker *Worker, err error){
	worker = new(Worker)
	err = worker.Init(ctx, wg)
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func (worker *Worker) Init(ctx context.Context, wg *sync.WaitGroup) (err error){
	endpoints := make([]string, 0)
	endpoints = append(endpoints, fmt.Sprintf("%s:%d", Config.Etcd.Host, Config.Etcd.Port))
	worker.Client, err = clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		DialTimeout: 5 *time.Second,
	})
	if err != nil {
		return err
	}

	worker.Ctx = ctx
	worker.WG = wg

	return nil
}

type Msg struct {
	MemoryTotal int64 `json:"mem_total"`
	MemoryUsed int64 `json:"mem_used"`
	MemoryFree int64 `json:"mem_free"`
	CPU int32 `json:"cpu"`
	Now time.Time `json:"now"`
	Tasks []string `json:"tasks"`
}

func (worker *Worker) Collect() (msg *Msg, err error) {
	msg = new(Msg)

	msg.CPU = 12
	msg.MemoryTotal = 17119559680
	msg.MemoryUsed = 10165350400
	msg.MemoryFree = 6954209280

	msg.Now = time.Now()

	for _, plugin := range(Config.Plugins) {
		msg.Tasks = append(msg.Tasks, plugin.Name)
	}

	return msg, nil
}

func (worker *Worker) Report() {
	timer := time.NewTimer(10 * time.Second)

	for {
		select {
			case <-timer.C:
				msg, _ := worker.Collect()
				value, err := json.Marshal(msg)
				if err != nil {
					log.Fatal(err)
				}
				key := fmt.Sprintf("/node/%s", Config.Node)

				worker.Client.Put(worker.Ctx, key, string(value))
				timer.Reset(10 * time.Second)

			case <-worker.Ctx.Done():
				log.Printf("reporter goroutine exit\n")
				worker.WG.Done()
				return
		}
	}
}

func (worker *Worker) Listener() {
	responses := worker.Client.Watch(worker.Ctx, fmt.Sprintf("/call/%s", Config.Node), clientv3.WithPrefix())
	for {
		select {
		case response := <-responses:
			for _, ev := range(response.Events) {
				log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if ev.Type == clientv3.EventTypePut {
					// 更新
					fmt.Println(ev)
				}else if ev.Type == clientv3.EventTypeDelete {
					// 删除
				}else{
					log.Printf("unknown event\n")
				}
				// 更新

			}
		case <-worker.Ctx.Done():
			log.Printf("listener goroutine exit\n")
			worker.WG.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)

	config, err := NewConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	worker, err := NewWorker(config, ctx, &wg)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(2)
	go worker.Report()
	go worker.Listener()

	select{
	case <-c:
		log.Printf("wait goroutine canceling\n")
		cancel()
		wg.Wait()
		log.Printf("all other goroutine canceled\n")
		log.Printf("main goroutine exit\n")
	}
}