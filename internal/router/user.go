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
		return err
	}
	c.JSON(http.StatusOK, user)
	h.logger.Info("Succesfully read user")
	return nil
}

func (h *Handler) createUser(c *gin.Context) error {
	var u types.User

	if err := c.BindJSON(&u); err != nil {
		return types.ErrBadRequest(fmt.Errorf("error binding JSON to struct: %w", err))
	}

	if err := h.userService.CreateUser(c, &u); err != nil {
		return err
	}
	h.logger.Info("Succesfully created user")
	return nil
}
