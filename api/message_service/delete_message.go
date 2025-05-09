package userserviceimpl
import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func (s *UserServiceImpl) DeleteMessageservice(c *gin.Context, messageID string, userID string) error {
	objId, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}
	err = s.messageAuth.DeleteMessagedb(c, objId, userID)
	if err != nil {
		return err
	}
	return nil
}
