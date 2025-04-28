# 留言系统接口文档

## 1. 提交留言

### 请求信息
- **接口URL**: `/api/v1/users/messages`
- **请求方法**: POST
- **Content-Type**: application/json

### 请求参数
```typescript
{
  content: string;      // 留言内容
  userId: string;       // 用户ID
  username: string;     // 用户名
  panoramaId: string;   // 全景图场景ID
}
```

### 响应格式
```typescript
{
  success: boolean;     // 请求是否成功
  data?: any;          // 可选的响应数据
  message?: string;    // 可选的响应消息
}
```

### 响应示例
```json
// 成功响应
{
  "success": true,
  "data": {
    // 返回的数据
  }
}

// 失败响应
{
  "success": false,
  "message": "留言提交失败"
}
```

## 2. 获取历史留言

### 请求信息
- **接口URL**: `/api/v1/users/messages`
- **请求方法**: GET
- **Content-Type**: application/json

### 请求参数
| 参数名 | 类型 | 位置 | 必填 | 说明 |
|--------|------|------|------|------|
| panoramaId | string | query | 是 | 全景图场景ID |

### 响应格式
```typescript
{
  success: boolean;     // 请求是否成功
  data?: Array<{       // 留言数据数组
    content: string;    // 留言内容
    // 其他可能的留言相关字段
  }>;
  message?: string;    // 可选的响应消息
}
```

### 响应示例
```json
{
  "success": true,
  "data": [
    {
      "content": "这是一条留言"
      // 其他留言相关字段
    }
  ]
}
```

## 前端处理说明

1. 提交留言时，前端会对留言内容进行空值校验
2. 提交成功后，会自动将新留言添加到弹幕显示列表中
3. 每条弹幕会随机生成以下属性：
   - top: 0-80 之间的随机值（控制垂直位置）
   - duration: 10-15 秒之间的随机值（控制动画持续时间）
   - color: 从预设颜色列表中随机选择

## 错误处理

1. 前端会处理以下错误情况：
   - 留言内容为空
   - 网络请求失败
   - 服务器响应错误

2. 错误提示使用 Element Plus 的 `ElMessage` 组件展示
   - 成功提示：`ElMessage.success`
   - 警告提示：`ElMessage.warning`
   - 错误提示：`ElMessage.error` 