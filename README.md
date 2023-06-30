# zsxq

## 功能介绍

1. 通过HTTP方式 获取知识星球 API，保存主题到数据库(MySQL)，并且可以导出为文件

## 使用方式

1. 执行 `cp config.example.yaml config.yaml` 生成配置文件
2. 电脑端登录知识星球，复制 cookie 到 `config.yaml` cookie 字段
3. 导入 MySQL (sql 文件在 `docs/initsql.sql`)
4. 执行 `make run` 即可运行该项目(适合 go 构建运行其他方式也可，ex [air](https://github.com/cosmtrek/air))

### 备注

1. 接口文档使用 swagger 生成，具体使用文档参考 [swaggo/swag](https://github.com/swaggo/swag)
2. 调试模式已打开，观察终端日志，如果接口报 `1059` 错误，过段时间重新执行接口即可。
3. 该项目之所以以 http 方式调用，而不是终端命令行是因为考虑到后期做管理后台.
