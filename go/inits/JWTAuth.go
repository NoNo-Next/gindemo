package inits

import (
	"errors"
	"gindemo/go/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 用户信息结构体
//type User struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义JWT的过期时间 设置2小时
const TokenExpireDuration = time.Hour * 2

// 定义一个secret
var JWTSecret = []byte("blue@yan")

// 1、生成JWT
func GenToken(username string) (string, error) {

	c := MyClaims{
		username, // 自定义的字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "jwt_test",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获取完整的编码后的字符串token
	return token.SignedString(JWTSecret)
}

// 2、解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 后面是一个匿名函数
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 3、定义一条 authHandler用于token的认证
func AuthHandler(c *gin.Context) {
	// 用户发送用户名与密码
	var user entity.User
	// 获取参数
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "无效的参数",
		})
		c.Abort()
		return
	}

	// 校验用户名与密码是否正确
	if user.UserName != "admin" || user.Password != "123456" {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "用户名或密码错误",
		})
		c.Abort()
		return
	}

	// 生成Token
	tokenString, _ := GenToken(user.UserName)
	c.JSON(200, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"token": tokenString, "tokenExpire": 7200},
	})
	return
}

// 4、实现校验的中间件
// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization") // 获取请求头中的数据
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8000")
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// *** 开始测试
/*func TestJWTAuth(t *testing.T) {
	// 默认引擎
	r := gin.Default()

	// 注册路由
	// 生成token的请求
	r.POST("/auth", authHandler)
	// home路由需要注册认证中间件
	r.GET("/home", JWTAuthMiddleware(), homeHandler)

	// 启动
	r.Run("127.0.0.1:9100")

}*/

// home路由
func HomeHandler(c *gin.Context) {
	// 获取参数
	username := c.MustGet("username").(string)

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}
