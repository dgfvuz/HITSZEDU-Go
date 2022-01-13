# hitszedu-go
go后端

*config*
写项目的配置文件。

*controller*
控制器层，验证提交的数据，将验证完成的数据传递给 service。

*service*
业务层，只完成业务逻辑的开发，不进行操作数据库。

*database*
数据库通用操作，不涉及业务代码。

*model*
数据库的ORM（进行数据库操作）。

*entity*
写返回数据的结构体。

*router*
写路由配置及路由的中间件（鉴权、日志、异常捕获）。

*util*
写项目通用工具类。

*static*
静态资源存放处。

*test*
测试用代码。

*session*
与session相关代码

*middleware*
中间件代码
