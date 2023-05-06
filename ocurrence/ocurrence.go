package ocurrence

type Occurrence struct {
	ID      int
	Color   string
	Nucleus int
	Tail    int
}

func NewOcurrence(id int, color string, nucleus int, tail int) *Occurrence {
	return &Occurrence{
		ID:      id,
		Color:   color,
		Nucleus: nucleus,
		Tail:    tail}
}
