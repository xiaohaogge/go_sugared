package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/routers/api"
)

// 获取全部房源信息
func GetRoomIndex(c *gin.Context) {
}

// 获取具体房源信息
func GetRoomDetail(c *gin.Context) {
	fmt.Println("file test add")
	getreq := &DeleteRoomReq{}
	if err := c.BindJSON(&getreq); err != nil {
		api.ApiResponse(c, 500, "request param error!!!")
	} else {
		fmt.Println("request pkg: ", getreq)
		err, res := FindOneRoomByPackageNumber(getreq.PackageNumber)
		if err != nil {
			fmt.Println("find err: ", err)
		}
		fmt.Println("res1: ", res)
	}

}

// 新增房源信息
func AddRoom(c *gin.Context) {
	fmt.Println("add room")
	roomAddReq := &Room{}
	if err := c.BindJSON(&roomAddReq); err != nil {
		api.ApiResponse(c, 400, "request param error....")
	} else {
		roomAddReq.ToPrint()
		if err := roomAddReq.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			api.ApiResponse(c, 500, "insert room error....")
		} else {
			api.ApiResponse(c, 200, "insert room sucess....")
		}
	}
}

// 编辑房源信息
func UpdateRoom(c *gin.Context) {
}

// 删除房源信息
func DeleteRoom(c *gin.Context) {
	fmt.Println("delete room")
	roomDeleteReq := &DeleteRoomReq{}
	if err := c.BindJSON(&roomDeleteReq); err != nil {
		api.ApiResponse(c, 400, "request param error....")
	} else {
		if err := roomDeleteReq.Delete(); err != nil {
			fmt.Println("delete err: ", err)
			api.ApiResponse(c, 500, "delete room error....")
		} else {
			api.ApiResponse(c, 200, "delete room sucess....")
		}
	}
}
