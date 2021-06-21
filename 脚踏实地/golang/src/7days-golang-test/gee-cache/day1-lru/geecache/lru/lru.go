package lru
/*
	LRU(least recently Used) 算法维护一个队列，如果某条记录被访问了,则移动到队尾，那么
	那么队首时最近最少访问的数据，淘汰该条记录即可. LRU(最近最少使用)相对于仅仅考虑时间因素和
	的FIFO和仅考虑访问频率的LFU，LRU算法可以认为时相对平衡的一种淘汰算法。
*/
import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.(并发访问是不安全的。)
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.(可选并在清除条目时执行)
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}




