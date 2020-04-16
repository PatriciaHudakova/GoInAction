//how channels can be used to monitor running times ==> background processed
package Ch7

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct { //runs a set of tasks before timeout
	interrupt chan os.Signal   //reports a signal from os
	complete  chan error       // complete channel reports processing is done
	timeout   <-chan time.Time //timeout reports time has been done
	tasks     []func(int)      //set of functions executed in index order
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner { //returns a ready to use runner
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) { // attaches tasks to runner
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error { //runs all tasks and monitors channel events
	signal.Notify(r.interrupt, os.Interrupt) //receive all interrupt based signals

	go func() { //run different tasks on different goroutines
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete: //signal when processing is done
		return err
	case <-r.timeout:
		return ErrTimeout //signal when time runs out
	}
}

func (r *Runner) run() error { //executes each registered task
	for id, task := range r.tasks {
		if r.gotInterrupt() { //check for interrupt signal from os
			return ErrInterrupt
		}
		task(id) //execute registered task
	}
	return nil
}

func (r *Runner) gotInterrupt() bool { //verifies interrupt signal
	select {
	case <-r.interrupt: //when interrupt event is sent
		signal.Stop(r.interrupt) //stop receiving any more signals
		return true
	default: //continue running as normal
		return false
	}
}
