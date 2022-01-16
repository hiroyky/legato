package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiroyky/legato/domain/errors"
	"github.com/hiroyky/legato/domain/request"
	"github.com/hiroyky/legato/infrastructure/database/repository"
	"github.com/hiroyky/legato/registry"
	"net/http"
	"path"
)

type musicController struct {
	trackRepo repository.TrackRepository
}

var MusicController *musicController = nil

func init() {
	MusicController = &musicController{
		trackRepo: registry.NewTrackRepository(),
	}
}

func (c *musicController) GetMusic(ctx *gin.Context) {
	var req request.GetMusicRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	track, err := c.trackRepo.GetByFilePathHash(ctx, req.PathHash)
	if err != nil {
		if errors.IsErrorCode(err, errors.TrackNotFoundError) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	http.ServeFile(ctx.Writer, ctx.Request, track.FilePath)
}

func (c *musicController) GetDownloadMusic(ctx *gin.Context) {
	var req request.GetDownloadMusicRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	track, err := c.trackRepo.GetByFilePathHash(ctx, req.PathHash)
	if err != nil {
		if errors.IsErrorCode(err, errors.TrackNotFoundError) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(track.FilePath)))
	http.ServeFile(ctx.Writer, ctx.Request, track.FilePath)
}
