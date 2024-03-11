# 间隙锁FAQ

## 间隙锁和 NextKey 锁的区别

todo...

## 查询条件是自增id，间隙锁要怎么锁啊？

todo...

## 间隙锁死锁如何解决

todo...

## X,GAP 是什么意思？

X代表互斥锁，GAP代表间隙锁，所以表明本次加的锁是间隙独占锁，合并起来就是X,GAP。

类似的还有:

- 记录共享和独占锁：S,REC_NOT_GAP & X,REC_NOT_GAP
- 意向共享和独占锁：IS & IX
- Next-Key共享和独占锁：S & X
- 插入意向独占锁：X,GAP,INSERT_INTENTION

通过查询 `performance_schema.data_locks` 我们能看到类似这样的输出：

```sql
$ select ENGINE_TRANSACTION_ID,EVENT_ID,INDEX_NAME,LOCK_TYPE,LOCK_MODE,LOCK_STATUS,LOCK_DATA from performance_schema.data_locks;
+-----------------------+----------+------------+-----------+-----------+-------------+-----------+
| ENGINE_TRANSACTION_ID | EVENT_ID | INDEX_NAME | LOCK_TYPE | LOCK_MODE | LOCK_STATUS | LOCK_DATA |
+-----------------------+----------+------------+-----------+-----------+-------------+-----------+
|                  1875 |       36 | NULL       | TABLE     | IX        | GRANTED     | NULL      |
|                  1875 |       36 | PRIMARY    | RECORD    | X,GAP     | GRANTED     | 10        |
+-----------------------+----------+------------+-----------+-----------+-------------+-----------+
```

## 字符串如何加间隙锁？

闭包表移动案例，如何加间隙锁。

TODO....

```sql
```
