package barrier

type Barrier struct {
	inChan  chan bool
	outChan chan bool
}

func NewBarrier(n int) *Barrier {
	bar := &Barrier{
		inChan:  make(chan bool),
		outChan: make(chan bool),
	}
	go bar.traverse(n)
	return bar
}

func (bar *Barrier) Wait() {
	bar.inChan <- true
	<-bar.outChan
}

func (bar *Barrier) traverse(n int) {

	for {
		for i := 0; i < n; i++ {
			<-bar.inChan
		}
		for i := 0; i < n; i++ {
			bar.outChan <- true
		}

	}
}
