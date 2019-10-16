package serializer

import (
	"eroauz/models"
)

type Volume struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	NovelID   uint   `json:"novel_id"`
	NovelName string `json:"novel_name"`
	FileID    uint   `json:"file_id"`
}

type VolumeResponse struct {
	Response
	Data Volume `json:"data"`
}

type VolumeListResponse struct {
	Response
	Count int      `json:"count"`
	All   int      `json:"all"`
	Data  []Volume `json:"data"`
	Next  bool     `json:"have_next"`
	Last  bool     `json:"have_last"`
	Pages int      `json:"pages"`
}

func BuildVolume(volume models.Volume) Volume {
	return Volume{
		ID:        volume.ID,
		Title:     volume.Title,
		Cover:     volume.Cover,
		NovelID:   volume.Novel.ID,
		NovelName: volume.Novel.Title,
		FileID:    volume.FileID,
	}
}

func BuildVolumeList(volumes []models.Volume) []Volume {
	var volumeList []Volume
	for _, a := range volumes {
		i := BuildVolume(a)
		volumeList = append(volumeList, i)
	}
	return volumeList
}

// BuildArchiveResponse 序列化文章响应
func BuildVolumeResponse(volumes models.Volume) VolumeResponse {
	return VolumeResponse{
		Data: BuildVolume(volumes),
	}
}

// BuildArchiveResponse 序列化文章列表响应
func BuildVolumeListResponse(volumes []models.Volume, all int, count int, next bool, last bool, pages int) VolumeListResponse {
	return VolumeListResponse{
		Count: count,
		All:   all,
		Data:  BuildVolumeList(volumes),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
