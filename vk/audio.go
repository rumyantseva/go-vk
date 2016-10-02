package vk

type AudioService struct {
	client *Client
}

type Tracks struct {
	Count int     `json:"response.count,omitempty"`
	Items []Track `json:"response.items,omitempty"`
}

type Track struct {
	ID       int64  `json:"id,omitempty"`
	OwnerID  int64  `json:"owner_id,omitempty"`
	Artist   string `json:"artist,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration int    `json:"count,omitempty"`
	Date     int    `json:"date,omitempty"`
	URL      string `json:"url,omitempty"`
	GenreID  int    `json:"genre_id,omitempty"`
}

func (s *AudioService) List(ownerID int64, count int, offset int) (*Tracks, *Response, error) {
	u := "audio.get"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var tracks = new(Tracks)
	resp, err := s.client.Do(req, tracks)
	if err != nil {
		return nil, resp, err
	}

	return tracks, resp, err
}
