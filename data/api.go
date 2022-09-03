package data

type Health struct {
	Health  string `json:"health" xml:"health" form:"health"`
	Message string `json:"message" xml:"message" form:"message"`
}

type Error struct {
	Code   string `json:"code" xml:"code" form:"code"`
	Reason string `json:"reason" xml:"reason" form:"reason"`
}

type RoomCreateRequest struct {
	Name string `json:"name" xml:"name" form:"name"`
}

type RoomCreate struct {
	Id int64 `json:"id" xml:"id" form:"id"`
}

type RoomConnectRequest struct {
	Name string `json:"name" xml:"name" form:"name"`
}

type Room struct {
	Id      int64    `json:"id" xml:"id" form:"id"`
	Name    string   `json:"name" xml:"name" form:"name"`
	Devices []Device `json:"devices" xml:"devices" form:"devices"`
}

type Device struct {
	Name      string `json:"name" xml:"name" form:"name"`
	Timestamp int64  `json:"timestamp" xml:"timestamp" form:"timestamp"`
}

type Status struct {
	Device    string `json:"device" xml:"device" form:"device"`
	Timestamp int64  `json:"timestamp" xml:"timestamp" form:"timestamp"`
	Type      string `json:"type" xml:"type" form:"type"`
	Name      string `json:"name" xml:"name" form:"name"`
	Mime      string `json:"mime" xml:"mime" form:"mime"`
	Hash      string `json:"hash" xml:"hash" form:"hash"`
}
