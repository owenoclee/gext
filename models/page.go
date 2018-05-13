package models

type Page struct {
	Threads []*Thread
}

func (p Page) IsEmpty() bool {
	return len(p.Threads) == 0
}

func (p Page) FirstThread() *Thread {
	if !p.IsEmpty() {
		return p.Threads[0]
	}
	return &Thread{}
}

func (p Page) LastThread() *Thread {
	if !p.IsEmpty() {
		return p.Threads[len(p.Threads)-1]
	}
	return &Thread{}
}

func (p Page) Board() string {
	if !p.IsEmpty() {
		return p.FirstThread().Board()
	}
	return ""
}
