# admin.claps.dev前端页面代码仓库

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

## 将前端部署到服务器端步骤

### 执行命令生成dist文件包
```
yarn build
```
即可将该文件包部署到服务器端的nginx或tomcat的相应目录下

## 相关端口及其他配置设置

### 端口配置
在vue.config.js文件中以下位置配置监听端口
```
"devServer": {
    "port": 8001,
  },
```

### 后端api配置
在utils/loginRequest.js以及request.js中配置后端api
```
baseURL: 'https://admin-api.claps.dev/',
```
