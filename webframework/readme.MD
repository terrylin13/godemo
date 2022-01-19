# WebFreamwork 框架
这是一个以学习一个web框架基本功能为导向的学习项目，了解框架核心为我们解决了什么问题。参考的是[7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee.html)。

## 重点学习内容

1. 实现标准 net/http 包里的 Handler 接口
> Handler用于响应一个HTTP request,接口方法ServerHTTP应该用来将response header和需要响应的数据写入到ResponseWriter中，然后返回。

2. 实现Context
> 设计上下文(Context),封装 Request 和 Response,提供对JSON、HTML 等返回类型的支持。Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。

3. 实现路由功能
> 一个web框架最核心的功能之一。简化各种访问URL的添加，将请求映射到函数及路由组的管理。这个框架将实现一个最简单的类似于tire树的节点树，来实现动态路由的功能。

4. 中间件


## 备注

本项目没有实现模板功能。主要是笔者主攻后端开发，现在都是前后端分离，且基于REST风格的API交互，故没有实现template功能。