package web

type MenuMasterResponse struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Path       *string `json:"path"`
	IconName   *string `json:"icon_name"`
	IsSubMenu  bool    `json:"is_submenu"`
	ParentName *string `json:"parent_name"`
	Create     bool    `json:"create"`
	Read       bool    `json:"read"`
	Update     bool    `json:"update"`
	Delete     bool    `json:"delete"`
}
