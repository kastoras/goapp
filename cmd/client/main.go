package main

import (
	"flag"
	"goapp/internal/pkg/watcher"
	"goapp/pkg/util"
	"log"
	"sync"
	"time"
)

func main() {
	connectionsNum := flag.Int("connections", 1, "Number of connections to create")
	flag.Parse()

	var wg sync.WaitGroup
	watchers := make([]*watcher.Watcher, *connectionsNum)
	for i := 0; i < *connectionsNum; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			w := watcher.New()
			watchers[id] = w

			if err := w.Start(); err != nil {
				log.Printf("Connection %d: failed to start: %v\n", id, err)
				return
			}
			runWatcher(w, id)
		}(i)
	}
	wg.Wait()
}

func runWatcher(w *watcher.Watcher, connNum int) {
	go func() {
		for counter := range w.Recv() {
			log.Printf("[conn #%d] : iteration: %d, value: %s\n",
				connNum, counter.Iteration, counter.Value)
		}
	}()

	for i := 0; ; i++ {
		rand := util.RandString(10)
		w.Send(rand)
		time.Sleep(1 * time.Second)
	}
}
