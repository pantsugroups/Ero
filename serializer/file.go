package serializer

import "eroauz/models"

type File struct {
	ID       uint   `json:"id"`
	FileName string `json:"filename"`
	Type     int    `json:"type"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
	Hash     string `json:"hash"`
	Token    string `token:"token"`
}
type FileResponse struct {
	Response
	Data File `json:"data"`
}

func BuildFile(file models.File) File {
	return File{
		ID:       file.ID,
		FileName: file.FileName,
		Type:     file.Type,
		UserID:   file.User.ID,
		UserName: file.User.Nickname,
	}
}
func BuildFileResponse(file models.File) FileResponse {
	return FileResponse{
		Data: BuildFile(file),
	}
}
func BuildDownloadFileResponse(file File) FileResponse {
	return FileResponse{
		Data: file,
	}
}
