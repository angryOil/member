# 필독 브랜치 규칙
main - prod, stage - stage , dev - 개발서버

pr 순서 feat => dev => (hotfix/bug)stage => main


1. 최신 dev브랜치에서 feature 만들기
2. dev에 push전 dev pull 받기
3. bug/hotfix 를 제외한 브랜치(ex:feat)로 main/stage에 직접pr금지

# member 서비스 입니다.

cafe Api에서 멤버 관련요청을합니다.

주기능은 cafe_id + user_id 를 통한 회원조회, 카페가입요청,요청 수락/거절 , 카페 벤 입니다.

## entity 구조
```text
member{ //cafe_id + user_id = uniq_key 
    id
    cafe_id
    user_id 
    nickname
    is_banned
    create_at
}
```

### makefile

```shell
# local postgres run (docker-compose)
make local-db
# local postgres migrate init
make local-init
# local postgres apply migrate
make local-migrate
```

# swagger 설정 [출처](https://www.soberkoder.com/swagger-go-api-swaggo/)

## dev 설정

```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

## main에

```code
   import (_ "[project명]/docs")
```

```shell
# swagger json 생성   swag init -g [project main path].go
swag init -g cmd/app/main.go
```

## [스웨거 링크](http://localhost:8082/swagger/index.html)