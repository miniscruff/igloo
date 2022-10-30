package mathf

type Sides struct {
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Top    float64 `json:"top"`
	Bottom float64 `json:"bottom"`
}

var (
	SidesZero = Sides{Left: 0, Right: 0, Top: 0, Bottom: 0}

	// stretches
	SidesStretchHorizontal = Sides{Left: 0, Right: 1, Top: 0, Bottom: 0}
	SidesStretchVertical   = Sides{Left: 0, Right: 0, Top: 0, Bottom: 1}
	SidesStretchBoth       = Sides{Left: 0, Right: 1, Top: 0, Bottom: 1}

	// different anchoring points
	SidesTopLeft      = SidesZero
	SidesTopRight     = Sides{Left: 1, Right: 1, Top: 0, Bottom: 0}
	SidesTopCenter    = Sides{Left: 0.5, Right: 0.5, Top: 0, Bottom: 0}
	SidesMiddleLeft   = Sides{Left: 0, Right: 0, Top: 0.5, Bottom: 0.5}
	SidesMiddleCenter = Sides{Left: 0.5, Right: 0.5, Top: 0.5, Bottom: 0.5}
	SidesMiddleRight  = Sides{Left: 1, Right: 1, Top: 0.5, Bottom: 0.5}
	SidesBottomLeft   = Sides{Left: 0, Right: 0, Top: 1, Bottom: 1}
	SidesBottomCenter = Sides{Left: 0.5, Right: 0.5, Top: 1, Bottom: 1}
	SidesBottomRight  = Sides{Left: 1, Right: 1, Top: 1, Bottom: 1}
)
