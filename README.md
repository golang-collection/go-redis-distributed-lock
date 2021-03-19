# go-redis-distributed-lock
go使用redis实现分布式锁

# 目录结构
- configs：用于存放配置相关文件
- global：用于存放全局设置
- internal/worker
    - lock：分布式锁实现
    - executor：正常的业务逻辑
    - main/main.go：work主函数
- pkg
    - cache：Redis pool连接
    - setting：与configs相对应读取配置文件