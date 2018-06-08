package user

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an user by the user identifier
// @Produce  json
// @Param id path int true "The user's database id index num"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id":3,"created_on":1516849721,"modified_on":0,"name":"3333","created_by":"4555","modified_by":"","state":0}],"total":29},"msg":"ok"}"
// @Router /v1/user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
