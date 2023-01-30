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

## 참고한 자료들
- [Redis를 통한 JWT Refresh Token 관리](https://sol-devlog.tistory.com/22)
- [Spring + Security + JWT + Redis를 통한 회원인증/허가 구현 (3) - 로그인 시 Access, Refresh Token 부여/ 사용](https://velog.io/@ehdrms2034/Spring-Security-JWT-Redis%EB%A5%BC-%ED%86%B5%ED%95%9C-%ED%9A%8C%EC%9B%90%EC%9D%B8%EC%A6%9D%ED%97%88%EA%B0%80-%EA%B5%AC%ED%98%84)