
//
//import (
//	"errors"
//	"fmt"
//	"sync"
//	"time"
//)
//
//const (
//	workerBits  uint8 = 10
//	numberBits  uint8 = 12
//	workerMax   int64 = -1 ^ (-1 << workerBits)
//	numberMax   int64 = -1 ^ (-1 << numberBits)
//	timeShift   uint8 = workerBits + numberBits
//	workerShift uint8 = numberBits
//	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
//)
//
//type Worker struct {
//	mu        sync.Mutex
//	timestamp int64
//	workerId  int64
//	number    int64
//}
//
//func NewWorker(workerId int64) (*Worker, error) {
//	if workerId < 0 || workerId > workerMax {
//		return nil, errors.New("Worker ID excess of quantity")
//	}
//	// 生成一个新节点
//	return &Worker{
//		timestamp: 0,
//		workerId:  workerId,
//		number:    0,
//	}, nil
//}
//
//func (w *Worker) GetId() int64 {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//	now := time.Now().UnixNano() / 1e6
//	if w.timestamp == now {
//		w.number++
//		if w.number > numberMax {
//			for now <= w.timestamp {
//				now = time.Now().UnixNano() / 1e6
//			}
//		}
//	} else {
//		w.number = 0
//		w.timestamp = now
//	}
//	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
//	return ID
//}
//func main() {
//	// 生成节点实例
//	node, err := NewWorker(1)
//	if err != nil {
//		panic(err)
//	}
//	for {
//		fmt.Println(node.GetId())
//	}
//}
package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func get() {

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()


	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())
}

func main(){
	go get()
}

