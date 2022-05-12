package album

import (
	"github.com/google/uuid"
)

// album represents data about a record album
type Album struct {
	ID		 string	 `json:"id"`
	Title	 string	 `json:"title"`
	Artist string  `json:"artist"`
	Price	 float64 `json:"price"`
}

func New(title string, artist string, price float64) Album {
	return Album{uuid.New().String(), title, artist, price}
}