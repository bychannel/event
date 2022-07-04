# Event 

轻量级的事件管理、调度工具库

已经实现的功能：
- 支持自定义事件
- 支持事件设置参数，并传递到执行方法
- 支持对一个事件添加多个监听器
- 支持设置事件监听器的优先级
- 支持事件名称使用"."进行分级，从而匹配一组事件
- 支持使用通配符 `*` 来监听全部事件的触发

## 主要方法

- `Listen(name string, listener Listener, priority ...int)` 注册事件监听
- `Subscribe(sbr Subscriber)`  订阅，支持注册多个事件监听
- `Publish(name string, params M) (error, Event)` 发布事件
- `MustPublish(name string, params M) Event`   发布事件，有错误则会panic
- `BatchPublish(es ...interface{}) (ers []error)` 一次发布多个事件
- `AsyncPublish(e Event)`   异步事件发布，使用协程

## 快速使用

见测试用例

## LICENSE

MIT LICENSE
