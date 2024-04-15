package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

func New() *Closer {
	closer := &Closer{}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		signal.Stop(ch)
		closer.CloseAll()
	}()

	return closer
}

func CloseAll() {
	globalCloser.CloseAll()
}

func Add(f func() error) {
	globalCloser.Add(f)
}

type Closer struct {
	f  []func() error
	mu sync.Mutex
}

func (c *Closer) Add(f func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.f = append(c.f, f)
}

func (c *Closer) CloseAll() {
	c.mu.Lock()
	f := c.f
	c.f = nil
	c.mu.Unlock()
	for _, fn := range f {
		err := fn()
		if err != nil {
			log.Println(err)
		}
	}
	os.Exit(0)
}
