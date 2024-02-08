package main

import (
	"fmt"
	"reflect"
)

func determinesTheTypeRuntime(value interface{}) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
		fmt.Println("It's Int:", reflect.TypeOf(value))
	case reflect.String:
		fmt.Println("It's String:", reflect.TypeOf(value))
	case reflect.Bool:
		fmt.Println("It's Bool:", reflect.TypeOf(value))
	case reflect.Chan:
		fmt.Println("It's Chan:", reflect.TypeOf(value))
	}
}

func main() {
	var (
		i  int
		s  string
		ch chan int
		b  bool
	)
	determinesTheTypeRuntime(i)
	determinesTheTypeRuntime(s)
	determinesTheTypeRuntime(ch)
	determinesTheTypeRuntime(b)
}
