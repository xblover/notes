# groupcache

## Summary

groupcache is a caching and cache-filling library, intended as a 
replacement for memcached in many cases.

For API docs and examples, see http://godoc.org/github.com/golang/groupcache

## Comparison to memcached

### **Like memcached**, groupcache:

* shards by key to select which peer is resonsible for that key
(按密钥分片以选择哪个对等方负责该密钥)

### **Unlike memcached**, groupcache:

* does not require running a separate set of servers, thus massively 
reducing deployment/configuration pain. groupcache is a client library 
as well as a server. It connects to its own peers.
(不需要运行一组单独的服务器，从而大大减少了部署/配置的痛苦。 groupcache 是一个客户端库，也是一个服务器。 它连接到它自己的同级。)

* 

## Loading process

## Users

## Presentations

## Help














