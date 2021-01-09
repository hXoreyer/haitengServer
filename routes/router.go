package routes

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"haitengServer/utils"
	"os"
)

//路由组注册
func userGroup(c *gin.Engine, Groupname string) *gin.RouterGroup {
	routeGroup := c.Group(Groupname)
	return routeGroup
}

//初始化
func InitRouter() {
	r := gin.Default()
	r.Use(utils.Cors())

	haiteng := userGroup(r, "haiteng") //注册haiteng路由组
	haitengRoute(haiteng)
	r.Run(":8050")
}

//海腾路由组路由
func haitengRoute(r *gin.RouterGroup) {
	r.POST("/login", loginPost)
	r.POST("/upload", getFile)
	r.GET("/download", downloadFile)
}

//haiteng路由组登录请求
func loginPost(c *gin.Context) {
	info := utils.Userinfo{}
	c.BindJSON(&info)

	if info.UserName == "lilijun" || info.UserName == "heminghai" {
		if info.UserPsw == "199836hkm" {
			c.JSON(200, gin.H{
				"info": "login",
				"key":  utils.GetKey(15),
			})
		} else {
			c.String(201, "%s 登陆失败，密码错误", info.UserName)
		}
	} else {
		c.String(202, "用户名错误")
	}

}

//haiteng路由组获取文件
func getFile(c *gin.Context) {
	file, _ := c.FormFile("filename")
	fileName := file.Filename
	fmt.Printf(fileName)
	c.SaveUploadedFile(file, fileName)
	str, err := utils.Execcmd("python autoExcel.py " + fileName)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "err",
			"err":    err,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "ok",
			"str":    str,
		})
	}
}

func downloadFile(c *gin.Context) {
	filename := "result.xlsx"
	// c.Query("filename")
	filepath := "./" + filename
	b := utils.PathExists(filepath)
	if b {
		file, err := os.Open(filepath)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		buf := bufio.NewWriter(file)
		buf.Flush()
		defer file.Close()
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File(filepath)
	}
}
