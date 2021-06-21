Go 在sync 包中提供了用于同步的一些基本原语，包括常见的sync.Mutex, sync.RWMutex, sync.WaitGroup, sync.Once, sync.Cond;

它们是一种相对原始的同步机制，在多数情况下，我们都应使用抽象层级更高的Channel 实现同步。



