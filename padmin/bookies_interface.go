package padmin

type Bookies interface {
	// AllBookies Gets raw information for all the bookies in the cluster
	AllBookies() (*AllBookiesResp, error)
	// ListRacksInfo Gets the rack placement information for all the bookies in the cluster
	ListRacksInfo() (*ListRacksInfoResp, error)
	// RemoveRacksInfo Removed the rack placement information for a specific bookie in the cluster
	RemoveRacksInfo(string) error
	// GetRacksInfo Gets the rack placement information for a specific bookie in the cluster
	GetRacksInfo(string) (*BookieInfo, error)
	// UpdateRacksInfo Updates the rack placement information for a specific bookie in the cluster
	// (note. bookie address format:`address:port`)
	UpdateRacksInfo(string, *BookieInfo) error
}

type RawBookieInfo struct {
	BookieId string `json:"bookieId"`
}

type BookieInfo struct {
	Rack     string `json:"rack"`
	Hostname string `json:"hostname"`
}

type AllBookiesResp struct {
	Bookies []RawBookieInfo `json:"bookies"`
}

type ListRacksInfoResp map[string]map[string]BookieInfo
