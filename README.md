# Aigouda 电商系统

这是一个基于 Vue.js 和 Go 开发的电商系统。

## 技术栈

### 前端
- Vue 3
- Ant Design Vue
- Vite
- Vue Router
- Pinia

### 后端
- Go
- Gin
- GORM
- MySQL

## 项目结构

```
.
├── frontEnd/          # 前端项目
│   ├── src/          # 源代码
│   ├── public/       # 静态资源
│   └── package.json  # 前端依赖配置
├── backend/          # 后端项目
│   ├── cmd/         # 主程序入口
│   ├── config/      # 配置文件
│   ├── internal/    # 内部包
│   └── go.mod       # Go 模块配置
└── README.md        # 项目说明
```

## 开发环境要求

- Node.js >= 16
- Go >= 1.16
- MySQL >= 5.7

## 本地开发

### 前端开发

```bash
cd frontEnd
npm install
npm run dev
```

### 后端开发

```bash
cd backend
go mod tidy
go run main.go
```

## 部署指南

### 1. 前端部署

```bash
cd frontEnd
npm install
npm run build
```

将 `dist` 目录下的文件部署到 Web 服务器。

### 2. 后端部署

```bash
cd backend
go build -o aigouda
./aigouda
```

### 3. 数据库配置

1. 创建数据库：
```sql
CREATE DATABASE aigouda CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改配置文件 `backend/config/config.yaml`：
```yaml
mysql:
  host: localhost
  port: 3306
  user: your_username
  password: your_password
  dbname: aigouda
```

## 注意事项

1. 确保服务器上已安装所需的环境
2. 配置文件中的敏感信息建议使用环境变量
3. 生产环境部署时注意修改 CORS 配置
4. 定期备份数据库 
