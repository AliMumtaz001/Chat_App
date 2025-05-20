package userserviceimpl

import (
	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserServiceImpl) UpdateMessageservice(c *gin.Context, messageID string, message models.Message) error {
	objId, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}
	err = s.messageAuth.UpdateMessagedb(c, objId, message)
	if err != nil {
		return err
	}
	return nil
}
