package main

//goroutine与线程
//可增长的栈
//OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB），一个goroutine的栈在其生命周期开始时只有
//很小的栈（典型情况下2KB），goroutine的栈是不固定的，他可以按需增大和缩小，goroutine的栈大小限制
//可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的goroutine也是可以的

// goroutine的调度
// GMP是Go语言运行时（runtime）层面的实现，是Go语言自己实现的一套调度系统。区别于操作系统调度OS线程
// G很好理解，就是goroutine，里面除了存放本goroutine信息外，还有与所在P的绑定等信息
//M是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是————映射的关系，一个goroutine最终是要放到M是执行的
//P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针、堆栈地址及地址边界），P会对自己管理的
//goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去
//全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务

func main() {

}
