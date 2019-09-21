package main

import (
	"codingInterview2/02_Singleton/example1"
	"codingInterview2/02_Singleton/example2"
	"codingInterview2/02_Singleton/example3"
	"codingInterview2/02_Singleton/example4"
	"codingInterview2/02_Singleton/example5"
	"sync"
)

func main() {
	w := sync.WaitGroup{}

	// 饿汉模式
	// 将包变量设置为私有，引入包时进行变量初始化，之后直接获取变量的值即可
	// 优点：1、并发安全 2、使用前预加载。
	// 缺点：1、只是引入却不使用也进行了预加载，造成内存浪费。
	example1.DoImport()
	example1.GetInstance().DoExample()

	// 懒汉模式
	// 不进行变量初始化，使用时再去初始化
	// 优点：1、用时加载，避免内存浪费
	// 缺点：1、不能使用前预加载 2、并发不安全（假如a、b两个协程同时运行GetInstance()，a协程通过了m==nil判断，准备运行m=&Manager{}。这时b协程也通过了m=nil判断。之后a、b协程都会运行m=&Manager{}。在并发场景下破坏了单例模式只能运行一次的特性。）
	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {
			example2.GetInstance().DoExample()
			w.Done()
		}()
	}
	w.Wait()

	// 懒汉加锁模式
	// 通过加互斥锁保证了并发安全。但是加锁造成了并发阻塞影响性能。
	// 优点：1、并发安全 2、用时加载，避免内存浪费
	// 缺点：1、不能使用前预加载 2、锁造成并发阻塞影响性能
	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {
			example3.GetInstance().DoExample()
			w.Done()
		}()
	}
	w.Wait()

	// 懒汉仅初始化时加锁模式
	// 加锁前先判断e是否已生成实例。这样只有e未初始化时才会加锁初始化，不用在每一次运行时都加锁。
	// 优点：1、并发安全 2、用时加载，避免内存浪费
	// 缺点：1、不能使用前预加载 2、代码繁琐
	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {
			example4.GetInstance().DoExample()
			w.Done()
		}()
	}
	w.Wait()

	// 懒汉仅初始化时加锁模式标准库版
	// 使用go标准库sync.Once的Do函数，保证单例模式。Do函数参数传递的函数，只会运行一次。
	// 优点：1、并发安全 2、用时加载，避免内存浪费 3、代码简洁
	// 缺点：1、不能使用前预加载
	// example5
	for i := 0; i < 2; i++ {
		w.Add(1)
		go func() {

			example5.GetInstance().DoExample()
			w.Done()
		}()
	}
	w.Wait()

}
