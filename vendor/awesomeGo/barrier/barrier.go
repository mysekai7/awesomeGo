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
	// fmt.Println("wait:\tbar.inChan <- true")
	bar.inChan <- true
	// fmt.Println("wait:\t<-bar.outChan")
	<-bar.outChan
}

func (bar *Barrier) traverse(n int) {
	for {
		for i := 0; i < n; i++ {
			//time.Sleep(time.Second * 2)
			// fmt.Println("traverse:\t<-bar.inChan")
			<-bar.inChan
		}
		for i := 0; i < n; i++ {
			//time.Sleep(time.Second * 2)
			// fmt.Println("traverse:\tbar.outChan <- true")
			bar.outChan <- true
		}
	}
}
