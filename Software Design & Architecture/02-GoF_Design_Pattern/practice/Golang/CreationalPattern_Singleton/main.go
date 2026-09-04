package main

import (
	"fmt"
	"sync"
)

var once = &sync.Once{}	// Lock dùng để bảo vệ việc tạo ra instance của singleton

type singleton struct{}

var singletonInstance *singleton

func getInstance() *singleton {
	if  singletonInstance == nil {
		once.Do(
			func() {
				fmt.Printf("Creating singleton instance\n")
				singletonInstance = &singleton{}
			},
		)
	} else {
		fmt.Printf("Singleton instance already created\n")
	}

	return singletonInstance
}

func main() {
	for i := 1; i< 30; i++ {
		go getInstance()
	}

	// Chờ cho tất cả các goroutine hoàn thành
	fmt.Scanln()
}