package utils

import sync "sync"

type Singleton struct {
	//
}

var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
