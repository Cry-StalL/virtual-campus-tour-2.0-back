# 留言 API 文档

## 接口端点

### 创建留言
在全景场景的特定位置创建新留言。

**URL:** `/api/v1/users/messages`  
**Method:** `POST`  
**Auth Required:** No (无需认证)

#### 请求体
```json
{
  "content": "string",      // 留言内容（最多50个字符）
  "userId": "int",       // 创建留言的用户ID
  "username": "string",     // 创建留言的用户名
  "panoramaId": "string",   // 全景场景ID
  "position": {
    "longitude": number,    // 全景图中的经度坐标
    "latitude": number      // 全景图中的纬度坐标
  }
}
```

#### 成功响应
- **Code:** 200 OK
- **Content:**
```json
{
  "success": true,
  "data": {
    "messageId": "string",    // 留言ID
    "content": "string",      // 留言内容
    "userId": "int",       // 用户ID
    "username": "string",     // 用户名
    "panoramaId": "string",   // 全景场景ID
    "position": {
      "longitude": number,    // 经度坐标
      "latitude": number      // 纬度坐标
    },
    "createdAt": "string"     // 创建时间
  }
}
```

#### 错误响应
- **Code:** 400 Bad Request
- **Content:**
```json
{
  "success": false,
  "message": "Error message"  // 错误信息
}
```

### 获取留言
获取特定全景场景的所有留言。

**URL:** `/api/v1/users/messages`  
**Method:** `GET`  
**Auth Required:** No (无需认证)

#### 查询参数
- `panoramaId` (required): 全景场景ID（必填）

#### 成功响应
- **Code:** 200 OK
- **Content:**
```json
{
  "success": true,
  "data": [
    {
      "messageId": "string",    // 留言ID
      "content": "string",      // 留言内容
      "userId": "int",       // 用户ID
      "username": "string",     // 用户名
      "panoramaId": "string",   // 全景场景ID
      "position": {
        "longitude": number,    // 经度坐标
        "latitude": number      // 纬度坐标
      },
      "createdAt": "string"     // 创建时间
    }
  ]
}
```

#### 错误响应
- **Code:** 400 Bad Request
- **Content:**
```json
{
  "success": false,
  "message": "Error message"  // 错误信息
}
```

## 注意事项
- 留言内容限制在50个字符以内
- 位置坐标使用全景图内的球面坐标系统，即经度和纬度
- 所有时间戳使用ISO 8601格式