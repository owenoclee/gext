package models

type Thread struct {
	Posts []Post
}

func (t Thread) IsEmpty() bool {
	return len(t.Posts) == 0
}

func (t Thread) HasReplies() bool {
	return len(t.Posts) > 1
}

func (t Thread) HasHiddenReplies() bool {
	return len(t.Posts) > 6
}

func (t Thread) OriginalPost() Post {
	if !t.IsEmpty() {
		return t.Posts[0]
	}
	return Post{}
}

func (t Thread) Replies() []Post {
	if t.HasReplies() {
		return t.Posts[1:]
	}
	return []Post{}
}

func (t Thread) HiddenReplies() []Post {
	if t.HasHiddenReplies() {
		replies := t.Replies()
		return replies[:len(replies)-5]
	}
	return []Post{}
}

func (t Thread) ShownReplies() []Post {
	if t.HasHiddenReplies() {
		return t.Posts[len(t.Posts)-5:]
	}
	return t.Replies()
}

func (t Thread) Id() uint32 {
	if !t.IsEmpty() {
		return t.OriginalPost().Id
	}
	return 0
}

func (t Thread) Board() string {
	if !t.IsEmpty() {
		return t.OriginalPost().Board
	}
	return ""
}

func (t Thread) Normalised() Thread {
	board := t.Board()
	for i := 0; i < len(t.Posts); i++ {
		t.Posts[i].Board = board
	}
	return t
}
