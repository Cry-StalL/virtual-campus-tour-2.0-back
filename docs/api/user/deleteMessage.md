### 删除留言
删除指定的留言。

**URL:** `/users/deleteMessage`  
**Method:** `POST`  
**Auth Required:** No (无需认证)

#### 请求体
```json
{
  "messageId": "string"    // 要删除的留言ID
}
```

#### 成功响应
- **Code:** 200 OK
- **Content:**
```json
{
  "code": 0,              // 0表示成功
  "message": "删除成功"
}
```

#### 错误响应
- **Code:** 400 Bad Request
- **Content:**
```json
{
  "code": 1,              // 非0表示失败
  "message": "Error message"  // 错误信息
}
```