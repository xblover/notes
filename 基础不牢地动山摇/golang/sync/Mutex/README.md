	## Mutex
	go语言的sync.Mutex由两个字段 state和 sema组成。
	
	```go
	type Mutex struct {
		state int32   // 当前互斥锁的状态
		sema  uint32  // 用于控制锁状态的信号量
	}
	```
	上述两个加起来只占8字节空间的结构体表示了go语言的互斥锁

	### 状态
	互斥锁的状态比较复杂，最低位分别表示mutexLocked, mutexWoken, mutexStarving ,剩下的位置用来表示当前有多少个goroutine 在等待互斥锁的释放
	
	```
		 waitersCount | starving | woken |locked 
	```
	在默认情况下，互斥锁的所有状态都是0，int32的不同位分别表示了不同的状态
	
	- mutexLocked, 表示互斥锁的锁定状态 
	- mutexWoken,  表示从正常模式被唤醒
	- mutexStarving ， 当前的互斥锁进入饥饿状态
	- waitersCount , 当前互斥锁上等待的goroutine 个数


	### 正常模式和饥饿模式
	sync.Mutex 有两种模式，正常模式和饥饿模式。

	//todo: 正常模式和饥饿模式都是什么以及它们有什么样的关系



	### 加锁和解锁
		互斥锁的加锁是靠sync.Mutex.Lock完成的

		源码即：当锁的状态是0时，将mutexLocked位置1
		```go
		func (m *Mutex) Lock() {
			if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
				return
			}
			m.lockSlow()
		}
	
		```
		如果互斥锁的状态不是0时就会调用sync.Mutex.lockSlow尝试通过自旋（Spinning）等方式等待锁的释放
			
		// todo: lockSlow 方法的详细解读
		```go
		func (m *Mutex) lockSlow() {
			var waitStartTime int64
			starving := false
			awoke := false
			iter := 0
			old := m.state
			for {
				if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
					if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 && atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
						awoken = true
					}
					runtime_doSpin()
					iter++
					old = m.state
					continue
				}
			}

			...
		}
		```







		互斥锁的解锁是靠sync.Mutex.Unlock完成的，

		源码如下：
		
		```go
		func (m *Mutex) Unlock() {

			// 先用sync/atomic.AddInt32 函数快速解锁
			new := atomic.AddInt32(&m.state, -mutexLocked)
			if new != 0 {
				// 如果函数返回状态不等于0，调用sync.Mutex.unlockSlow开始慢速解锁
				m.unlockSlow(new)
			 }
			 
			 //如果函数返回的状态=0，当前goroutine 就成功解锁了互斥锁
		}

		```

		// todo: sync.Mutex.unlockSlow 源码的详细解读

		```go
		func (m *Mutex) unlockSlow(new int32) {
			if (new+mutexLocked)&mutexLocked == 0 {
				throw("sync: unlock of unlocked mutex")
			}
			...
		}
		```

		### 总结
		互斥锁的加锁过程比较复杂，它涉及自旋，信号量以及调度等概念

		//todo :加锁的详细总结


		互斥锁的解锁相对简单一些

		//todo: 解锁的详细总结


		












