package seriallzer

type UserUpdateReq struct {
	Stature    int    `json:"stature"`
	Weight     int    `json:"weight"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
}
