package resource

import videoV1 "video/api/qvbilam/video/v1"

type DramaResource struct {
	Id          int64              `json:"id"`
	Name        string             `json:"name"`
	Introduce   string             `json:"introduce"`
	Icon        string             `json:"icon"`
	User        *UserResource      `json:"user"`
	Episode     *[]EpisodeResource `json:"episode"`
	CreatedTime int64              `json:"created_time"`
}

func (r *DramaResource) Collection(response *videoV1.DramaListResponse) *[]DramaResource {
	var dramas []DramaResource
	for _, drama := range response.Drama {
		var episodes []EpisodeResource
		for _, e := range drama.Episode {
			episodes = append(episodes, EpisodeResource{
				Id:      e.Id,
				Episode: e.Episode,
				Video: VideoResource{
					Id:        e.Video.Id,
					FileId:    e.Video.FileId,
					Name:      e.Video.Name,
					Introduce: e.Video.Introduce,
					Icon:      e.Video.Icon,
				},
			})
		}

		dramas = append(dramas, DramaResource{
			Id:          drama.Id,
			Name:        drama.Name,
			Introduce:   drama.Introduce,
			Icon:        drama.Icon,
			User:        nil,
			Episode:     &episodes,
			CreatedTime: drama.CreatedTime,
		})
	}

	return &dramas
}
