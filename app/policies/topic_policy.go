// Package policies 用户授权
package policies

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/pkg/auth"
)

// CanModifyTopic 判断当前用户是否有权限修改
func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
