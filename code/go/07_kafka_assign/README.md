# kafka assign 模式

启动producer：

```bash
$ go run main.go consumer.go ddd
```

启动多个consumer:

```bash
$ go run main.go consumer.go
```

可以看到，多个consumer都消费到了生产者生产的消息！