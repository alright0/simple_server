package user_handler

import (
	"main/internal/dto"
	"main/internal/repositories/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteUser(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData dto.DeleteUserRequest

		if err := c.ShouldBindJSON(&userData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := c.Request.Context()

		isSuccess, err := users.SoftDeleteUser(ctx, pool, userData.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !isSuccess {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User delete failed"})
			return
		}

		c.AbortWithStatus(http.StatusNoContent)
	}
}
