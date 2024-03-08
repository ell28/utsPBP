package controller

import (
	m "UTS/model"
	"encoding/json"
	"net/http"
)

func ShowAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	defer rows.Close()

	rooms := []m.Room{}
	for rows.Next() {
		room := m.Room{}
		err := rows.Scan(&room.ID, &room.Name, &room.GameID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		rooms = append(rooms, room)
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func ShowDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	roomID := r.URL.Query().Get("id")

	row := db.QueryRow("SELECT * FROM rooms WHERE id = ?", roomID)
	room := m.Room{}
	err := row.Scan(&room.ID, &room.Name)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rows, err := db.Query(`
        SELECT r.id, r.name, p.id, p.id_account, a.username
        FROM participants p
        JOIN accounts a ON p.id_account = a.id
		JOIN rooms r ON p.id_room = r.id
        WHERE p.id_room = ?`, roomID)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	defer rows.Close()

	participants := []m.Participant{}
	for rows.Next() {
		participant := m.Participant{}
		err := rows.Scan(&participant.ID, &participant.Account.ID, &participant.Account.Username)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		participants = append(participants, participant)
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomResponse
	response.Status = 200
	response.Data = struct {
		Room         m.Room        `json:"room"`
		Participants []Participant `json:"participants"`
	}{
		Room:         room,
		Participants: participants,
	}
	json.NewEncoder(w).Encode(response)
}
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	roomID := r.URL.Query().Get("room_id")
	accountID := r.URL.Query().Get("account_id")
	var response m.InsertResponse

	row := db.QueryRow("SELECT G.max_player FROM Games G JOIN Rooms R ON G.id = R.id_game WHERE R.id = ? ", roomID)
	var maxPlayer int
	err := row.Scan(&maxPlayer)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	countRow := db.QueryRow("SELECT COUNT(*) FROM participants WHERE id_room = ?", roomID)
	var participantCount int
	err = countRow.Scan(&participantCount)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if participantCount >= maxPlayer {
		response.Status = 400
		response.Message = "Room is full"
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response.Status = 200
	response.Message = "Participant inserted to room successfully"
	json.NewEncoder(w).Encode(response)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	roomID := r.URL.Query().Get("room_id")
	accountID := r.URL.Query().Get("account_id")
	var response m.InsertResponse

	_, err := db.Exec("DELETE FROM participants WHERE id_room = ? AND id_account = ?", roomID, accountID)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response.Status = 200
	response.Message = "Participant left the room successfully"
	json.NewEncoder(w).Encode(response)
}
