package data

type Health struct {
	Health  string
	Message string
}

type Error struct {
	Code   string
	Reason string
}

type RoomCreateRequest struct {
	Name string
}

type RoomCreate struct {
	Id int64
}

type RoomConnectRequest struct {
	Name string
}

type Room struct {
	Id      int64
	Name    string
	Devices []Device
}

type Device struct {
	Name      string
	Timestamp int64
}

type Status struct {
	Device    string
	Timestamp int64
	Type      string
	Name      string
	Mime      string
	Hash      string
}
