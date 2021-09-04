package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/legato/domain/errors"
	"github.com/legato/domain/request"
	"github.com/legato/infrastructure/database/repository"
	"github.com/legato/registry"
	"io/ioutil"
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

	buf, err := ioutil.ReadFile(track.FilePath)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(track.FilePath)))
	ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf)))
	ctx.Writer.Write(buf)
}
