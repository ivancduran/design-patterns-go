package singleton

import (
	"sync"
)

type singleton struct {
	Name string
	ID   int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
