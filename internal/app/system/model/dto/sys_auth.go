package dto

type MenuMeta struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	IsLink      string `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
}

type UserMenu struct {
	Id        uint   `json:"id"`
	Pid       uint   `json:"pid"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Path      string `json:"path"`
	*MenuMeta `json:"meta"`
}

type UserMenus struct {
	*UserMenu `json:""`
	Children  []*UserMenus `json:"children"`
}
