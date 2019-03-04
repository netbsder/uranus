## uranus项目说明

### 1 概述

uranus是一个web快速开发脚手架，旨在为用户提供开箱即用的常见web系统的基本服务，帮助用户提高开发效率。

它包括两个部分：

- uranus

  uranus后端服务，主要对外提高基于Restful API 服务接口，使用go语言开发。

- uranus-web-manage

  uranus后端服务的管理前端界面，基于VUE框架开发。

![uranus项目](https://upload-images.jianshu.io/upload_images/15874442-0792ec87e9cdba1d.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 1.1 uranus-service

uranus-service内的服务被拆分成多个服务模块，被拆分的每个模块既可以单独运行，作为web服务使用，也可以非常方便的嵌入到其它项目中。同时所有服务模块是无状态的，一方面便于水平扩展，另一方面便于向微服务架构迁移。

计划包含服务有：

- uranus-service-actuator - 应用监控管理服务
- uranus-service-ucms - 统一用户管理服务
- uranus-service-sms - 短信服务
- uranus-service-oss - 对象存储服务
- uranus-service-ems - 邮件服务
- uranus-service-wechat - 微信接口服务

#### 1.2 uranus-web-manage

详见：https://github.com/netbsder/uranus-web-manage

### 2. 使用技术&框架

- [GIN](https://github.com/gin-gonic/gin)
- [XORM](https://github.com/go-xorm/xorm)
- 待完善...

### 3. uranus简单部署

![uranus简单部署](https://upload-images.jianshu.io/upload_images/15874442-ef1b50f0a5015680.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)