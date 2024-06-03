package model

type ExoPlanet struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistanceFromEarth float64 `json:"distance_from_earth"`
	Radius            float64 `json:"radius"`
	Mass              float64 `json:"mass,omitempty"`
	Type              string  `json:"type"`
}
