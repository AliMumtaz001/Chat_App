package userserviceimpl
import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func (s *UserServiceImpl) DeleteMessageService(c *gin.Context, messageID string) error {
	objId, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}
	err = s.messageAuth.DeleteMessagedb(c, objId)
	if err != nil {
		return err
	}
	return nil
}
