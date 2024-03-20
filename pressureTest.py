import requests

def test_get_user_info():
    url = "http://127.0.0.1:8090/users/info"
    headers = {
        "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTA5NDMyNTEsInVzZXJfaWQiOjF9.u1Bfw8e7in6V8J2sNABrcXGPGv4kzHX3a9AYPb0j7rQ"
    }

    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        print("GET /users/info successful")
        print(response.json())
    else:
        print(f"GET /users/info failed with status code: {response.status_code}")
        print(response.text)

# 调用函数进行测试
for i in range(100000):
    test_get_user_info()