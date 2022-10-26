package resource

type DramaResource struct {
	Id          int64              `json:"id"`
	Name        string             `json:"name"`
	Introduce   string             `json:"introduce"`
	Icon        string             `json:"icon"`
	User        *UserResource      `json:"user"`
	Episode     *[]EpisodeResource `json:"episode"`
	CreatedTime int64              `json:"created_time"`
}
