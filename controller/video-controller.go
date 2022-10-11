package controller

import (
	"gintoki/entity"
	"gintoki/service"
	"gintoki/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

type VideoController interface {
	FindAll() ([]entity.Video, error)
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

// FindAll godoc
// @Summary Get all videos
// @Description Get all videos
// @Tags videos
// @Accept json
// @Produce json
// @Param video body entity.Video true "Get all videos"
// @Success 200 {array} entity.Video
// @Router /api/videos [GET]
func (c *controller) FindAll() ([]entity.Video, error) {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	return c.service.Save(video)
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, parseErr := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if parseErr != nil {
		return parseErr
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	return c.service.Update(video)
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	return c.service.Delete(video)
}
