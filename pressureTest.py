import requests

def test_get_user_info():
    # 定义测试获取用户信息的函数
    url = "http://127.0.0.1:8090/users/info"
    headers = {
        "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTA5NDMyNTEsInVzZXJfaWQiOjF9.u1Bfw8e7in6V8J2sNABrcXGPGv4kzHX3a9AYPb0j7rQ"
    }

    # 发送 GET 请求获取用户信息
    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        # 若响应状态码为 200，打印成功信息及响应数据
        print("GET /users/info successful")
        print(response.json())
    else:
        # 若响应状态码不为 200，打印失败信息及响应内容
        print(f"GET /users/info failed with status code: {response.status_code}")
        print(response.text)

# 调用函数进行测试
for i in range(100000):
    test_get_user_info()
