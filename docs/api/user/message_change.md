```json
{
  "code": 0,        // 状态码，0表示成功，其他值表示失败
  "message": "success", // 操作结果描述
  "data": {}        // 返回的数据，可能为空
}
```

## 错误码说明

| 错误码 | 说明 |
| ----- | ---- |
| 0     | 成功 |
| 400   | 请求参数错误 |
| 1001  | 用户不存在 |
| 1002  | 用户名已存在/密码加密失败 |
| 1003  | 邮箱格式不正确 |
| 1004  | 登录尝试次数过多 |
| 2001  | 邮箱已被注册 |
| 2002  | 邮箱格式不正确 |
| 2003  | 发送过于频繁 |
| 2004  | IP发送次数过多 |
| 2005  | 邮件发送失败 |

## API 端点

### 1. 更新用户名

修改用户的用户名。

**请求**

- 方法: `POST`
- URL: `/api/v1/users/updateUsername`
- 内容类型: `application/json`

**请求参数**

| 字段     | 类型   | 必填 | 描述     |
| -------- | ------ | ---- | -------- |
| userId   | uint   | 是   | 用户ID   |
| username | string | 是   | 新用户名，长度4-20个字符 |

**请求示例**

```json
{
  "userId": 1,
  "username": "newUsername"
}
```

**成功响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "username": "newUsername",
    "updateTime": "2024-03-20T10:30:45Z"
  }
}
```

**错误响应**

```json
{
  "code": 1001,
  "message": "用户不存在",
  "data": null
}
```

或

```json
{
  "code": 1002,
  "message": "用户名已存在",
  "data": null
}
```

### 2. 重置密码

重置用户的登录密码。

**请求**

- 方法: `POST`
- URL: `/api/v1/users/resetPassword`
- 内容类型: `application/json`

**请求参数**

| 字段     | 类型   | 必填 | 描述     |
| -------- | ------ | ---- | -------- |
| userId   | uint   | 是   | 用户ID   |
| password | string | 是   | 新密码，长度6-20个字符 |

**请求示例**

```json
{
  "userId": 1,
  "password": "newSecurePassword"
}
```

**成功响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "updateTime": "2024-03-20T10:35:22Z"
  }
}
```

**错误响应**

```json
{
  "code": 1001,
  "message": "用户不存在",
  "data": null
}
```

或

```json
{
  "code": 1002,
  "message": "密码加密失败",
  "data": null
}
```