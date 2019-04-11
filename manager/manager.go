package manager

import (
	"fmt"
	"sync"
)

var m *Manager
var once sync.Once

func SharedInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

type Manager struct {
	Cache string
}

func (p *Manager) Save(value string) {
	fmt.Println("set cache...", value)
	p.Cache = value
}

func (p *Manager) Verify(str string) {
	fmt.Printf("%s say: %s\n", str, p.Cache)
}
