package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Node  struct {
	MemoryTotal int64 `json:"mem_total"`
	MemoryUsed int64 `json:"mem_used"`
	MemoryFree int64 `json:"mem_free"`
	CPU int32 `json:"cpu"`
	Now time.Time `json:"now"`
	Tasks []string `json:"tasks"`
}

type Core struct {
	Client *clientv3.Client
	Ctx context.Context
	WG *sync.WaitGroup

	// node infos
	Nodes map[string]*Node

	// task map nodes
	Tasks map[string][]string
}

func NewCore(Config *config, ctx context.Context, wg *sync.WaitGroup) (core *Core, err error){
	core = new(Core)
	err = core.Init(ctx, wg)
	if err != nil {
		return nil, err
	}
	return core, nil
}

func (core *Core) Init(ctx context.Context, wg *sync.WaitGroup) (err error) {
	endpoints := make([]string, 0)
	endpoints = append(endpoints, fmt.Sprintf("%s:%d", Config.Etcd.Host, Config.Etcd.Port))

	core.Client, err = clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}

	core.Ctx = ctx
	core.WG = wg

	core.Nodes = make(map[string]*Node)
	core.Tasks = make(map[string][]string, 0)

	return nil
}

func (core *Core) Close() {
	core.Client.Close()
}

func (core *Core) UpdateTask(node *Node, msg *Node) {
	// task in msg, but task not in node, add
	for _, task := range(msg.Tasks) {
		flag := false
		for i, _ := range(node.Tasks) {
			if task == node.Tasks[i] {
				flag = true
				break
			}
		}
		if !flag {
			// add
			node.Tasks = append(node.Tasks, task)
			fmt.Printf("add task %s\n", task)
			// new task
		}
	}

	// task in node, but task not in msg, delete
	for _, task := range(node.Tasks) {
		flag := false
		for i, _ := range(msg.Tasks) {
			if task == msg.Tasks[i] {
				flag = true
				break
			}
		}
		if !flag {
			// del
			fmt.Printf("del task %s\n", task)
		}
	}
}

// watch /node/*
func (core *Core) WatchNode() (err error){
	responses := core.Client.Watch(core.Ctx, "/node", clientv3.WithPrefix())
	for {
		select {
		case response := <-responses:
			for _, ev := range(response.Events){
				if ev.Type == clientv3.EventTypePut {
					// 更新
					fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					var msg Node
					nodeName := string(ev.Kv.Key[6:])
					json.Unmarshal(ev.Kv.Value, &msg)

					_, ok := core.Nodes[nodeName]
					if !ok {
						core.Nodes[nodeName] = new(Node)
						core.Nodes[nodeName].Tasks = make([]string, 0)
					}
					node := core.Nodes[nodeName]
					node.Now = msg.Now

					core.UpdateTask(node, &msg)
				}else if ev.Type == clientv3.EventTypeDelete {
					// 删除
				} else {
					fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				}
			}
		case <-core.Ctx.Done():
			fmt.Printf("core goroutine exit()\n")
			core.WG.Done()
			return nil
		}
	}
}

func (core *Core) FindNode(task string) (node string){
	return "node1"
}

func (core *Core) Scheduler() (err error){
	responses := core.Client.Watch(core.Ctx, "/rpc", clientv3.WithPrefix())
	for {
		select {
		case response := <-responses:
			for _, ev := range(response.Events) {
				if ev.Type == clientv3.EventTypePut {
					names := strings.Split(string(ev.Kv.Key), "/")
					node := core.FindNode(names[2])
					k := fmt.Sprintf("/call/%s_%s/%s/request", names[2], node, names[3])
					core.Client.Put(core.Ctx, k, string(ev.Kv.Value))
					core.Client.Delete(core.Ctx, string(ev.Kv.Key))
				}else if ev.Type == clientv3.EventTypeDelete {
					// delete
					fmt.Printf("delete successfully\n")
				}else {
					fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				}
			}
		case <-core.Ctx.Done():
			fmt.Printf("scheduler goroutine exit()\n")
			core.WG.Done()
			return nil
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

	core, err := NewCore(config, ctx, &wg)
	if err != nil {
		log.Fatal(err)
	}
	defer core.Close()

	wg.Add(2)

	go core.WatchNode()
	go core.Scheduler()

	select {
	case <-c:
		log.Printf("wait goroutine canceling\n")
		cancel()
		wg.Wait()
		log.Printf("all other goroutine canceled\n")
		log.Printf("main goroutine exit\n")
	}
}
