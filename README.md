# Getting Started with Create React App
## BackEnd Run 
1. BackEnd src 디렉토리 이동.
``` shell
PS C:\Users\ckdqj\go\MyShoppingMall> cd .\backend\
PS C:\Users\ckdqj\go\MyShoppingMall\backend> cd .\src\
```
2. 실행
``` shell 
PS C:\Users\ckdqj\go\MyShoppingMall\backend\src> go run main.go
2023/01/07 23:24:57 Main log....
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /products                 --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.GetProducts-fm (3 handlers)
[GIN-debug] GET    /promos                   --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.GetPromos-fm (3 handlers)  
[GIN-debug] POST   /user/:id/signout         --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.SignOut-fm (3 handlers)    
[GIN-debug] GET    /user/:id/orders          --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.GetOrders-fm (3 handlers)  
[GIN-debug] POST   /users/charge             --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.Charge-fm (3 handlers)     
[GIN-debug] POST   /users/signin             --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.SignIn-fm (3 handlers)     
[GIN-debug] POST   /users                    --> github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest.HandlerInterface.AddUser-fm (3 handlers)    
[GIN-debug] Listening and serving HTTP on localhost:8000
```

## FrontEnd Run
1. 
