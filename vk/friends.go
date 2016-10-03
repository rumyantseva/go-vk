package vk

type FriendsService struct {
	client *Client
}

type Friends struct {
	Count int     `json:"response.count,omitempty"`
	Items []Friend `json:"response.items,omitempty"`
}

type Friend struct {
	ID       int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
}

func (s *FriendsService) Get(userID int64, count int, offset int) (*Tracks, *Response, error) {
	u := "friends.get"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var friends = new(Friends)
	resp, err := s.client.Do(req, friends)
	if err != nil {
		return nil, resp, err
	}

	return friends, resp, err
}
