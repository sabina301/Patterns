package code

import (
	"fmt"
	"sync"
)

var singletonInstance *singleton

type singleton struct {
}

var lock = &sync.Mutex{}

func GetInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			singletonInstance = &singleton{}
			fmt.Println("CREATE")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
	return singletonInstance
}
