### 获取用户留言

**请求URL:** `/api/v1/users/getUserMessages`

**请求方式:** GET

**请求参数:**

| 参数名 | 类型   | 必填 | 描述   |
| ------ | ------  | ---- | ------ |
| userId | unsigned big int | 是   | 用户ID |

**响应示例:**

```json
{
  "code": 0,
  "message": "获取成功",
  "data": [
    {
      "id": 12,
      "userId": 5,
      "content": "图书馆的环境真的很好，学习氛围浓厚!",
      "location": "图书馆",
      "createTime": "2023-12-25 14:30:00"
    },
    {
      "id": 8,
      "userId": 5,
      "content": "教学楼空调温度太低了，冬天都快冻僵了",
      "location": "第一教学楼",
      "createTime": "2023-12-20 09:15:00"
    }
  ]
}
```
这是个人界面获取历史留言的API接口说明，后端看看