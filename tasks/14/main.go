package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func determinesTheTypeRuntime(value interface{}) {
	/*
		Благодаря функции TypeOf из библиотеки reflect мы можем определить тип переменной во время выполнения, а не во время компиляции.
		Если interface == nil, то TypeOf возвращает nil
	*/
	refType := reflect.TypeOf(value)

	// Если применить метод Kind() к nil, то мы получим панику (invalid memory address or nil pointer dereference)
	if refType == nil {
		return
	}

	//	Используя метод Kind(), мы можем получить базовый тип переменной
	switch refType.Kind() {
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
	determinesTheTypeRuntime(nil)
}
