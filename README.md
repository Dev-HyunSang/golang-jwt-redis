# golang-jwt-rieds
Go언어와 JWT(JSON Web Token)의 취약점을 Redis로 보완하는 프로젝트입니다.

## 프로젝트 소개
본 프로젝트는 기존의 JWT의 보안의 취약점을 보완하고자 Redis를 사용해서 취약점을 보완하는 프로젝트입니다.  
프로젝트는 [20221229-JWT에서 Redis를 사용하는 이유](https://hyunsang.dev/TIL/Golang/20221229-JWT%EC%97%90%EC%84%9C-Redis%EB%A5%BC-%EC%82%AC%EC%9A%A9%ED%95%98%EB%8A%94-%EC%9D%B4%EC%9C%A0.html)에 대한 궁금증에서부터 시작되었습니다.

## 작동 방식
### 회원가입
1. 사용자가 회원가입을 요청합니다.
2. 이메일과 패스워드를 받습니다.
   - 평문 패스워드를 암호화 시킵니다. 
3. 모든 정보를 데이터베이스에 기록합니다.  
데이터베이스에 평문 패스워드는 기록하지 않습니다.  
`bcrypt`로 암호화된 패스워드를 기록합니다.
4. 정상적으로 완료 되었으면 사용자가에 회원가입 완료 응답을 보냅니다.

### 로그인
1. 사용자가 로그인을 요청합니다.
2. 사용자가 로그인을 요청한 정보가 올바른지 데이터베이스 저장된 정보와 검증합니다.
3. 올바른 사용자가 맞다면
   1. Redis 기록된 Access Token(접근에 관여하는 토큰)와  Refresh Token(재발급에 관여하는 토큰)를 비교합니다.

## REST API Docs
### POST `/auth/join`
#### Request
```json
{
    "email": "me@hyunsang.dev",
    "password": "q1w2e3r4",
    "nick_name": "HyunSang Park"
}
```

#### Response
```json
{
    "meta": {
        "status": "error",
        "status_code": 200,
        "success": true,
        "message": "성공적으로 요청하신 유저를 만들었어요!"
    },
    "data": {
        "id": 1,
        "user_uuid": "a1ba7027-cb01-41dd-9416-352aeea0ca99",
        "user_email": "me@hyunsang.dev",
        "user_password": "$2a$10$q/4bNIrzUrO.N7m8jYLyvudFuP21Ek4Kv2Fm/SkNDZY1Xt0URWZKi",
        "user_nickname": "HyunSang Park",
        "created_at": "2023-01-30T15:49:51.91299+09:00",
        "updated_at": "2023-01-30T15:49:51.91299+09:00"
    },
    "responsed_at": "2023-01-30T15:49:51.9261+09:00"
}
```

### POST `/auth/login`
#### Request
```json
{
    "email": "me@hyunsang.dev",
    "password": "q1w2e3r4"
}
```

#### Response
```json
{
    "meta": {
        "status": "success",
        "status_code": 200,
        "success": true,
        "message": "성공적으로 로그인을 완료했어요!"
    },
    "data": {
        "user_uuid": "a1ba7027-cb01-41dd-9416-352aeea0ca99",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjYyZjM5OWUyLWM3ZTctNGZmMi1hNTY0LWUxMTQ5Y2NiOWYxNyIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY3NTA2MjQ2OSwidXNlcl91dWlkIjoiYTFiYTcwMjctY2IwMS00MWRkLTk0MTYtMzUyYWVlYTBjYTk5In0.IxURueOzBaBbRJt7q-O26rXtH8g2_6dCXS9JiEz73kY",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU2NjYzNjksInJlZnJlc2hfdXVpZCI6Ijc1MjI2MmQ4LWU0ZjUtNDg5Zi05YjFmLTBiZDIyZmRhYjIyZSIsInVzZXJfdXVpZCI6ImExYmE3MDI3LWNiMDEtNDFkZC05NDE2LTM1MmFlZWEwY2E5OSJ9.gS_RH65_JUUYcPULj5tE-q36cG2guBzkQvcEHCpOEK8"
    },
    "responsed_at": "2023-01-30T15:52:49.567428+09:00"
}
```
![](./assets/login-redis-01.png)

![](./assets/login-redis-02.png)

### POST `/auth/logout`
#### Request
**Header:**  
```text
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImNkNGJhMTcyLTRhZWMtNGVmYi1hNWNmLTZiZWVmZjVjYzk5OCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY3NTA2NjMzNiwidXNlcl91dWlkIjoiOTkzNjBmNzktMWJlZi00NzdmLTg2MzgtZWU0MTI3ZGZhYjE3In0.G5iW0m2SL4mMHZR13TAy-7It4gDAeVidcYoC5fc1vwc
```

#### Response
```json
{
    "meta": {
        "status": "ok",
        "status_code": 200,
        "success": true,
        "message": "성공적으로 로그아웃 되셨습니다."
    },
    "data": 1,
    "responsed_at": "2023-01-31T18:47:45.112174+09:00"
}
```

### POST `/todo/create`
#### Request
**Header:**  
```text
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImNkNGJhMTcyLTRhZWMtNGVmYi1hNWNmLTZiZWVmZjVjYzk5OCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY3NTA2NjMzNiwidXNlcl91dWlkIjoiOTkzNjBmNzktMWJlZi00NzdmLTg2MzgtZWU0MTI3ZGZhYjE3In0.G5iW0m2SL4mMHZR13TAy-7It4gDAeVidcYoC5fc1vwc
```

**Body:**  
```json
{
    "todo_title": "Hello, World!",
    "todo_context": "Hello, World!"
}
```

#### Response
```json
{
    "meta": {
        "status": "ok",
        "status_code": 200,
        "success": true,
        "message": "성공적으로 새로운 할일을 생성했습니다."
    },
    "data": {
        "id": 4,
        "todo_uuid": "439c9657-1c25-41ee-a45e-90bbef0307d2",
        "user_uuid": "99360f79-1bef-477f-8638-ee4127dfab17",
        "todo_title": "Hello, World!",
        "todo_context": "Hello, World!",
        "updated_at": "2023-01-30T17:00:41.94033+09:00",
        "crated_at": "2023-01-30T17:00:41.94033+09:00"
    },
    "responsed_at": "2023-01-30T17:00:41.951056+09:00"
}
```

### POST `/todo/update`
#### Request
**Header:**  
```text
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImNkNGJhMTcyLTRhZWMtNGVmYi1hNWNmLTZiZWVmZjVjYzk5OCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY3NTA2NjMzNiwidXNlcl91dWlkIjoiOTkzNjBmNzktMWJlZi00NzdmLTg2MzgtZWU0MTI3ZGZhYjE3In0.G5iW0m2SL4mMHZR13TAy-7It4gDAeVidcYoC5fc1vwc
```

**Body:**
```json
{
    "todo_uuid": "0004bbf4-233f-4b03-93ac-e9728224acb1",
    "todo_title": "Hello, World!2",
    "todo_context": "Hello, World!2"
}
```

#### Response
```json
{
    "meta": {
        "status": "ok",
        "status_code": 200,
        "success": true,
        "message": "성공적으로 할일을 수정했어요!"
    },
    "data": 1,
    "responsed_at": "2023-01-31T18:52:23.817192+09:00"
}
```

## 참고한 자료들
- [Redis를 통한 JWT Refresh Token 관리](https://sol-devlog.tistory.com/22)
- [Spring + Security + JWT + Redis를 통한 회원인증/허가 구현 (3) - 로그인 시 Access, Refresh Token 부여/ 사용](https://velog.io/@ehdrms2034/Spring-Security-JWT-Redis%EB%A5%BC-%ED%86%B5%ED%95%9C-%ED%9A%8C%EC%9B%90%EC%9D%B8%EC%A6%9D%ED%97%88%EA%B0%80-%EA%B5%AC%ED%98%84)