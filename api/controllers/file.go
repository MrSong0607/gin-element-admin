package controllers

import (
	"admin/utils"
	"net/url"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type File struct{}

func init() {
	_router = append(_router, File{})
}

func (f File) RouterRegist(r *gin.Engine) {
	g := r.Group("file")
	{
		g.POST("upload", ApiWrap(f.Upload))
	}
}

func (File) Upload(c *gin.Context) (any, error) {
	file, e := c.FormFile("file")
	if e != nil {
		return nil, e
	}

	fileName := utils.Hex(utils.GetUUID()) + filepath.Ext(file.Filename)
	e = c.SaveUploadedFile(file, path.Join(utils.GetConfig().FileServerDir, fileName))

	u, _ := url.Parse(utils.GetConfig().FileServerBaseUrl)
	u.Path = path.Join(u.Path, fileName)
	return u.String(), e
}
