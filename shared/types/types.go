package types

type (
	Route struct {
		Distance float64     `json:"distance"`
		Duration float64     `json:"duration"`
		Geometry []*Geometry `json:"geometry"`
	}

	Geometry struct {
		Coordinates []*Coordinate `json:"coordinates"`
	}

	Coordinate struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)
