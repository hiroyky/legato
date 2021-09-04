package request

type GetMusicRequest struct {
	PathHash string `uri:"path_hash" binding:"required"`
}

type GetDownloadMusicRequest struct {
	PathHash string `uri:"path_hash" binding:"required"`
}
