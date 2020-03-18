package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/models"
)

// 获取全部房源信息
func GetRooms(c *gin.Context) {
}

// 获取具体房源信息
func GetRoomDetail(c *gin.Context) {
	fmt.Println("file test add")
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

// 新增房源信息
func AddRooms(c *gin.Context) {
	fmt.Println("add room")
	roomAddReq := &models.Room{}
	if err := c.BindJSON(&roomAddReq); err != nil {
		data := &models.RoomBaseResponse{
			Code: 500,
			Msg:  "xingbuxingo "}
		c.JSON(500, data)
		// context.JSON(200, gin.H{
		// 	"code": 500,
		// 	"msg":  "roomadd err",
		// },
		// )
	} else {
		roomAddReq.ToPrint()
		models.InsertRoom(roomAddReq)
		data2 := &models.RoomBaseResponse{
			Code: 200,
			Msg:  "序列化成功!!!! "}
		c.JSON(200, data2)

		// return &model.RoomBaseResponse{
		// 	Code: 200,
		// 	Msg:  "xingbuxingo ",
		// }
	}
}

// 编辑房源信息
func UpdateRooms(c *gin.Context) {
}

// 删除房源信息
func DeleteRooms(c *gin.Context) {
}
