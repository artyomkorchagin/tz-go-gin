package router

import (
	"fmt"
	"net/http"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) error {
	id := c.Params.ByName("id")
	user, err := h.userService.ReadUser(c, id)
	if err != nil {
		return fmt.Errorf("error reading user from DB: %w", err)
	}
	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) createUser(c *gin.Context) error {
	var u types.User

	if err := c.BindJSON(&u); err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("error binding JSON to struct: %w", err),
		}
	}

	if err := h.userService.CreateUser(c, &u); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error inserting user into DB: %v", err),
		}
	}
	return nil
}
