config 写项目的配置文件。

controller 控制器层，验证提交的数据，将验证完成的数据传递给 service。

service 业务层，只完成业务逻辑的开发，不进行操作数据库。

repository 数据库操作层，比如写，多表插入，多表查询等，不写业务代码。

model 数据库的ORM。

entity 写返回数据的结构体 && 写 controller 层方法参数验证的结构体。

router 写路由配置及路由的中间件（鉴权、日志、异常捕获）。

util 写项目通用工具类。
