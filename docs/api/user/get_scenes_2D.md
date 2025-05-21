### 获取用户留言

**请求URL:** `/api/v1/scenes/getScenes2D`

**请求方式:** GET

**请求参数:**

| 参数名 | 类型   | 必填 | 描述   |
| ------ | ------ | ---- | ------ |


**响应示例:**

```json
{
  "code": 0,
  "message": "获取成功",
  "data": [
    {
      "sceneId": 1,
      "sceneName": "图书馆",
      "scene_2D_URL": "assets/panos/streets/scene_2D/1.png"
    },
    {
      "sceneId": 24,
      "sceneName": "天琴中心",
      "scene_2D_URL": "assets/panos/streets/scene_2D/24.png"
    }
  ]
}
```
这是获取场景选择的2D图片