## uranus项目说明

### 1 概述

uranus项目旨在为用户提供开箱即用的常见web系统的基本服务，帮助用户提高开发效率。它包括两个子项目，uranus-web-manage和uranus-service。

![uranus项目](https://upload-images.jianshu.io/upload_images/15874442-0792ec87e9cdba1d.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

在本系列文章中，我们将通过该项目来阐述web后端开发的过程及细节。

### 1.1 uranus-service

通用web后端服务接口，包括统一用户管理服务(ucms)，短信服务(sms)，对象存储服务(oss)、邮件服务(ems)、微信接口服务(wechat)。uranus-service内的服务拆分成多个服务模块，并且所有服务模块是无状态的，一方面便于水平扩展，另一方面便于后续向微服务架构迁移。

#### 1.2 uranus-web-manage

uranus中后台管理系统界面

...待完善

### 2 uranus简单部署

![uranus简单部署](https://upload-images.jianshu.io/upload_images/15874442-ef1b50f0a5015680.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)