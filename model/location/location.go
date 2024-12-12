package location

type Location struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp float64 `json:"timestamp"`
	Address   string  `json:"address"`
}

func (l *Location) TimestampInt() int64 {
	return int64(l.Timestamp)
}
