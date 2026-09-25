# Authenticate JWT với Golang

Cấu trúc thư mục:
```bash
golang_study/
├── app/
│   ├── http/
│   ├── model/
│   ├── repository/
│   └── service/
├── config/
├── helper/
├── routes/
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── index.go
└── README.md
```

-   Trước tiên là mình sẽ tạo một helper.go trong thư mục helper để custom response:
```go
package helper

import "strings"

type Response struct {
    Status  bool        `json:"status"`
    Message string      `json:"message"`
    Errors  interface{} `json:"errors"`
    Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
    res := Response{
        Status:  status,
        Message: message,
        Errors:  nil,
        Data:    data,
    }
    return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
    splittedError := strings.Split(err, "\n")
    res := Response{
        Status:  false,
        Message: message,
        Errors:  splittedError,
        Data:    data,
    }
    return res
}
```

-   Tại đây thường run sẽ gọi lệnh  `go run index.go`  để start server vậy xem trong file index.go này (sử dụng framework  [gin](https://github.com/gin-gonic/gin)  cho bài hướng dẫn này)

```go
package main

import (
    "golang_api/routes"
    "gorm.io/gorm"
    "golang_api/config"
)

var (
    db             *gorm.DB                  = config.SetupDatabaseConnection()
)

func main() {
    defer config.CloseDatabaseConnection(db)
    router := routes.InitRouter()
    router.Run()
}
```

Như code thì bạn thấy file này sẽ để gọi router và connect database,

-   Setup database trong file config.go tại thư mục config
```go
package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "golang_api/app/model"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
    errEnv := godotenv.Load()
    if errEnv != nil {
        panic("Failed to load env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to create a connection to database")
    }
    db.AutoMigrate(&model.Book{}, &model.User{})
    return db
}

// Đóng kết nối
func CloseDatabaseConnection(db *gorm.DB) {
    dbSQL, err := db.DB()
    if err != nil {
        panic("Failed to close connection from database")
    }
    dbSQL.Close()
}
```

-   Trong folder  `routes`  tạo thêm một file  `index.go`
```go
package routes

import (
    "golang_api/app/http/controller"
    "golang_api/app/http/middleware"
    "golang_api/app/service"
    "golang_api/app/repository"
    "github.com/gin-gonic/gin"
    "golang_api/config"
    "gorm.io/gorm"
)

var (
    db             *gorm.DB                  = config.SetupDatabaseConnection()
    userRepository repository.UserRepository = repository.NewUserRepository(db)
    jwtService     service.JWTService        = service.NewJWTService()
    authService    service.AuthService       = service.NewAuthService(userRepository)
    authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func InitRouter() *gin.Engine {
    routes := gin.Default()

    authRoutes := routes.Group("api/auth")
    {
        authRoutes.POST("/register", authController.Register)
        authRoutes.POST("/login", authController.Login)
    }

userRoutes := routes.Group("api/user", middleware.AuthorizeJWT(jwtService)) 
{
	 userRoutes.GET("/profile", userController.Profile) 
	 userRoutes.PUT("/profile", userController.Update)  
}
    

    return routes
}
```  

Ở đây tạo ra 2 api đó là `login`, `register` để giành cho việc authenticate. tiếp tục xem `authController sẽ xử lý gì nhé`. Như khởi tạo bạn có thể tháy mình đã sử dụng cả `service`, `repository`... cho nên cùng xem lần lượt các instance này làm gi tiếp nhé
```go
package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "golang_api/app/http/request"
    "golang_api/app/model"
    "golang_api/helper"
    "golang_api/app/service"
)

type AuthController interface {
    Register(ctx *gin.Context)
}

type authController struct {
    authService service.AuthService
    jwtService  service.JWTServic
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
    return &authController{
        authService: authService,
        jwtService:  jwtService,
    }
}

func (c *authController) Register(ctx *gin.Context) {
    var registerRequest request.RegisterRequest
    errRequest := ctx.ShouldBind(&registerRequest)
    if errRequest != nil {
        response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }
}

func (c *authController) Login(ctx *gin.Context) {
    var loginRequest request.LoginRequest
    errRequest := ctx.ShouldBind(&loginRequest)
    if errRequest != nil {
        response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }
    authResult := c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
    if v, ok := authResult.(model.User); ok {
        generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
        v.Token = generatedToken
        response := helper.BuildResponse(true, "OK!", v)
        ctx.JSON(http.StatusOK, response)
        return
    }
    response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
    ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *userController) Profile(context *gin.Context) {
    authHeader := context.GetHeader("Authorization")
    token, err := c.jwtService.ValidateToken(authHeader)
    if err != nil {
        panic(err.Error())
    }
    claims := token.Claims.(jwt.MapClaims)
    id := fmt.Sprintf("%v", claims["user_id"])
    user := c.userService.Profile(id)
    res := helper.BuildResponse(true, "OK", user)
    context.JSON(http.StatusOK, res)
}

```
Như trong `routes` mình đã import thì ở trong controller này mình đã gọi `server`, `repository` vào để sử dụng. Trước tiên là thực hiện function `Register`, việc đầu tiên là chúng ta phải validate, và tại đây có tạo một file request.go trong thư mục `http/request`

```go
package request
type RegisterRequest struct {
    Name     string `json:"name" form:"name" binding:"required"`
    Email    string `json:"email" form:"email" binding:"required,email" `
    Password string `json:"password" form:"password" binding:"required"`
}
```

Tiếp theo sau khi pass validate thì vào bước xử lý nhé, ở đây mình có check duplicate mail sau đó thì sẽ tạo tài khoản
```go
...
if !c.authService.IsDuplicateEmail(registerRequest.Email) {
    response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
    ctx.JSON(http.StatusConflict, response)
} else {
        createdUser := c.authService.CreateUser(regist erRequest)
    token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
    createdUser.Token = token
    response := helper.BuildResponse(true, "OK!", createdUser)
    ctx.JSON(http.StatusCreated, response)
}
```

Hàm để check duplicate mail, tại:

authService.go
```go
func (service *authService) IsDuplicateEmail(email string) bool {
    res := service.userRepository.IsDuplicateEmail(email)
    return !(res.Error == nil)
}


func (j *jwtService) GenerateToken(UserID string) string {
    claims := &jwtCustomClaim{
        UserID,
        jwt.StandardClaims{
            ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
            Issuer:    j.issuer,
            IssuedAt:  time.Now().Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(j.secretKey))
    if err != nil {
        panic(err)
    }
    return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
    return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
        if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
        }
        return []byte(j.secretKey), nil
    })
}

``` 

userRepository.go
```go
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
    var user model.User
    return db.connection.Where("email = ?", email).Take(&user)
}

``` 

Bây gio tiến hành test nhé với cổng mặc định là 8080 Đây là trường hợp không pass validate


### Authenticated

Trong folder  `app/http`  bạn tạo môt một folder middleware có chứa file  `jwt.auth-.go`
```go
package middleware

import (
    "log"
    "net/http"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang_api/helper"
    "golang_api/app/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
            c.AbortWithStatusJSON(http.StatusBadRequest, response)
            return
        }
        token, err := jwtService.ValidateToken(authHeader)
        if token.Valid {
            claims := token.Claims.(jwt.MapClaims)
            log.Println("Claim[user_id]: ", claims["user_id"])
            log.Println("Claim[issuer] :", claims["issuer"])
        } else {
            log.Println(err)
            response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
            c.AbortWithStatusJSON(http.StatusUnauthorized, response)
        }
    }
}

```

<!--stackedit_data:
eyJoaXN0b3J5IjpbLTEyMDYwMDYzNTFdfQ==
-->