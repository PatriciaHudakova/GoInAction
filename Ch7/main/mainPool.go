package main

import (
	"github.com/PatriciaHudakova/GoInAction/Ch7"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct { //simulates resource
	ID int32
}

//resource release management
func (dbConn *dbConnection) Close() error {
	log.Println("close: connection", dbConn.ID)
	return nil
}

var idCounter int32 //gives each connection a unique id

func createConnection() (io.Closer, error) { //factory method
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	//create pool
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	//perform queries from pool
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()

	//close pool
	log.Println("shutdown program")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	//acquire connection from pool
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn) //release conncetion back to the pool

	//wait to simulate query response
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
