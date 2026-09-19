package note

import "time"

type Note struct {
	ID        int64
	Text      string
	CreatedAt time.Time
}
