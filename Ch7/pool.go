package Ch7

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct { //manages a set of resources that can be shared safely
	m         sync.Mutex
	resources chan io.Closer            //buffered channel
	factory   func() (io.Closer, error) //function type
	closed    bool                      //flag
}

var ErrPoolClosed = errors.New("Pool has been closed")

//creates a pool that manages resources
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) { //retrieves resource from pool
	select {
	case r, ok := <-p.resources: //check for a free resource
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default: //provide new resource if none available
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) { //places new resource onto the pool
	p.m.Lock()
	defer p.m.Unlock()

	//if pool closed, discard resources
	if p.closed {
		r.Close()
		return
	}

	select {
	//attempt to place new resource onto queue
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	//if queue already at capacity, close resource
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

//shutdown pool and close all existing resources
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	//if pool closed, do nothing
	if p.closed {
		return
	}

	//set pool as closed
	p.closed = true

	//close channel
	close(p.resources)

	//close resources
	for r := range p.resources {
		r.Close()
	}
}
