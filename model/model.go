package model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"max_players"`
}

type Room struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	GameID int    `json:"game_id"`
}

type Participant struct {
	ID      int     `json:"id"`
	Room    Room    `json:"room"`
	Account Account `json:"account"`
}

type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type RoomDetail struct {
	Room         Room          `json:"room"`
	Participants []Participant `json:"participants"`
}

type RoomResponse struct {
	Status int `json:"status"`
	Data   struct {
		Room struct {
			ID           int           `json:"id"`
			RoomName     string        `json:"room_name"`
			Participants []Participant `json:"participants"`
		} `json:"room"`
	} `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type InsertResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
