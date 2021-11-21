package entity

// Ad -.
type Ad struct {
	AdType      string `json:"ad_type" example:"banner"`
	Description string `json:"description"   example:"kupi skyrim"`
	AdText      string `json:"text"   example:"kupi skyrim"`
}
