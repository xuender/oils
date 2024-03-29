
1.0.51 / 2022-12-13
==================

* feat: logs 改用 logrus
* feat: syncs pool
* feat: add poster
* feat: 协程池
* refactor: remove maps, slices
* feat: base add const
* doc: event bus
* chore: update require
* test: 延迟消息总线

1.0.50 / 2022-12-09
==================

* feat: EventBus
* feat: go style ob
* feat: log console
* fix: RoutineGroup Add order error

1.0.49 / 2022-12-09
==================

* feat: remove map, slice等
* test: 更换 assert

0.5.48 / 2022-12-04
==================

* feat: log console

0.5.47 / 2022-11-30
==================

* fix: RoutineGroup Add order error
* test: times
* refactor: yield
* feat: syncs.RoutineGroup 增加 Inc 方法
* refactor: 使用 lo

0.5.46 / 2022-11-23
==================

* fix: 限频无法打满

0.5.45 / 2022-11-22
==================

* feat: InstanceNum 取消订阅

0.5.44 / 2022-11-22
==================

* feat: 队列限频优化

0.5.43 / 2022-11-21
==================

* test: queue.SetSize
* feat: 基于队列的限频

0.5.41 / 2022-11-21
==================

* feat: 队列及限频队列
* feat: oss 唯一ID
* feat: maths 最大最小
* feat: 生成序列
* doc: InstanceNum
* feat: syncs.Limits 增加 Keys

0.5.40 / 2022-11-10
==================

* refactor: 删除 distributed
* refactor: add tags
* doc: Tag 索引
* feat: dbs.Tag
* doc: treemap.Interator
* test: times.Debounced
* test: syncs/maps 并发测试

0.4.39 / 2022-11-10
==================

* feat: TreeMap 和 Maps 修改key规则
* refactor: rename
* test: TreeMap.Add

0.4.38 / 2022-11-09
==================

* feat: 线程安全的左倾红黑树Map
* feat: 多管道合并
* feat: treemap
* feat: TreeMap

0.4.37 / 2022-11-03
==================

* fix: 内存限频锁异常

0.4.36 / 2022-11-03
==================

* fix: 限频计时器有问题修复

0.4.35 / 2022-11-03
==================

* fix: 滑动窗口限频计时器失效错误
* test: 实验代码
* test: 实验代码
* refactor: 循环重构

0.4.34 / 2022-11-03
==================

* feat: redis滑动窗口限频优化
* test: example 分离
* style: 式样修改

0.4.33 / 2022-11-02
==================

* fix: 滑动窗口限频错误

0.4.32 / 2022-11-02
==================

* feat: 滑动窗口限频
* rdb limit 并发修改

0.4.31 / 2022-11-02
==================

* feat: redis 分布式并发限制

0.4.30 / 2022-10-31
==================

* fix: Counter 并发错误修复

0.4.29 / 2022-10-31
==================

* feat: 高速计数器

0.4.28 / 2022-10-26
==================

* feat: logs

0.4.26 / 2022-10-08
==================

* style: lint
* feat: 防抖函数

0.3.22 / 2022-07-22
==================

* feat: 消费方法
* feat: 字符串通配符匹配
* feat: 字符串匹配
* feat: 编辑距离
* feat: 生产消费接口
* feat: 消费协程号
* feat: add syncs.Consumer 多协程消费
* feat: 切片删除不改变原始切片
* feat: 切片替换不改变原始切片
* feat: add Sample, Shuffle
* refactor: 废弃重复的API
* refactor: 废弃重复的API
* refactor: slice Del, DelAll, Join, Unique

0.3.21 / 2022-07-06
==================

* feat: add Sub, Chunk, String2Map, Tag2Map
* fix: concurrent map writes

0.3.20 / 2022-07-05
==================

* test: session
* feat: dbs add Cache, BitMap
* feat: gins tests

0.2.19 / 2022-06-27
==================

* feat: add ios.NewWriter
* feat: sqls.Params
* feat: sqls
* feat: sqls
* Merge remote-tracking branch 'refs/remotes/origin/dev' into dev
* test: distributed_mock
* doc: 分布式
* feat: add SliceMap

0.2.18 / 2022-06-18
==================

* feat: 分布式锁

0.2.17 / 2022-06-18
==================

* feat: add bases.Values
* feat: 多异常合并
* refactor: base.Conversion
* refactor: base.Conversion
* doc: ParsePass ParseMust
* feat: add ParsePass & ParseMust
* feat: max, min
* chore: gomonkey/v2 v2.7.0

0.2.16 / 2022-06-14
==================

* feat: add base.Errors

0.2.15 / 2022-06-13
==================

* feat: base.Recover
* Merge remote-tracking branch 'refs/remotes/origin/dev' into dev
* feat: add times
* feat: add dbs.id
* feat: add oss.IsDir, oss.IsFile

0.2.14 / 2022-06-08
==================

* doc: must, pass, sorts, parse
* fix: ios.ContainsReader
* test: sytle
* test: add nets tests

0.2.13 / 2022-06-02
==================

* test: base.Split
* feat: add base.Split
* feat: add nets.Dirs
* chore: workflows

0.2.12 / 2022-05-30
==================

* test: add tests
* test: add tests
* doc: Deprecated

0.2.11 / 2022-05-28
==================

* feat: add default FileMode
* feat: add oss.Exist
* feat: add oss.Exec
* test: gomonkey
* feat: Panic

0.2.10 / 2022-05-26
==================

* style: index
* refactor: Panic 改成 Must
* feat: base.Sorts
* style: test
* feat: add Append

0.1.9 / 2022-05-22
==================

* chore: v0.1.9
* feat: i18n
* feat: 最大值,最小值,优化漏桶
* feat: add LeakyBucket 漏桶算法
* doc: doc

0.1.8 / 2022-05-17
==================

* feat: cryptos 加密
* feat: add Filter

0.1.7 / 2022-05-12
==================

* feat: hashs
* chore: go1.18

0.1.6 / 2022-05-11
==================

* feat: gin & gorm
* feat: add oss.Signal

0.1.5 / 2022-05-03
==================

* feat: nets.Service
* refactor: slice
* refactor: unique
* refactor: slice

0.1.4 / 2022-04-30
==================

* feat: add Unique

0.1.3 / 2022-04-29
==================

* feat: add ios, logs, oss
* fix: float to bytes and bytes to float
* test: pass
* chore: coverage
* doc: codecov

0.0.2 / 2022-04-29
==================

* chore: typecheck
* chore: typecheck
* chore: typecheck
* chore: typecheck
