package code

type Adapter struct {
	w *Win
}

func NewAdapter(w *Win) *Adapter {
	return &Adapter{w}
}

func (a *Adapter) InsertInSquarePORT() {
	a.w.InsertInCirclePORT()
}
