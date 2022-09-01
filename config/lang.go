package config

type Strings struct {
	Lang    string
	Console ConsoleStrings
	API     APIStrings
}

type ConsoleStrings struct {
	Starting      string
	SnowFlakeNode string
	HandlerExit   string
	Error         ConsoleErrorStrings
}

type ConsoleErrorStrings struct {
	DBOpen              string
	DBNotSupported      string
	SnowFlake           string
	DeadCode            string
	DuplicatedID        string
	FiberResponse       string
	StorageOpen         string
	StorageNotSupported string
}

type APIStrings struct {
	Error APIErrorStrings
}

type APIErrorStrings struct {
	EmptyRoomName      string
	EmptyDeviceName    string
	DupDevice          string
	DeviceNotConnected string
	BadHash            string
	ContentMismatch    string
	MimeMismatch       string
}

var (
	Msg map[string]Strings
)

func loadMsg() {
	Msg = map[string]Strings{}
	Msg["ja-JP"] = Strings{
		Lang: "japanese",
		Console: ConsoleStrings{
			Starting:      "Octane API Serverを開始しています...",
			SnowFlakeNode: "このプロセスのSnowFlake Nodeを作成しました。",
			HandlerExit:   "ハンドラーが終了しました。",
			Error: ConsoleErrorStrings{
				DBOpen:              "DBとの接続中に問題が発生しました。",
				DBNotSupported:      "その種類のDBはサポートされていません。",
				SnowFlake:           "Snowflake Nodeを作成中にエラーが発生しました。",
				DeadCode:            "到達してはならないコードパスに到達しました。",
				DuplicatedID:        "IDの重複が検出されました。SnowFlakeのNodeが重複していないかどうか確認することをおすすめします。",
				FiberResponse:       "Fiberが要求に返答中にエラーが発生しました。",
				StorageOpen:         "ストレージを開いている途中にエラーが発生しました。",
				StorageNotSupported: "その種類のストレージはサポートされていません。",
			},
		},
		API: APIStrings{
			APIErrorStrings{
				EmptyRoomName:      "部屋名が空です。",
				EmptyDeviceName:    "デバイス名が空です。",
				DupDevice:          "デバイス名が重複しています。",
				DeviceNotConnected: "そのようなデバイスは接続されていません。",
				BadHash:            "ハッシュの形式が不正です。",
				ContentMismatch:    "コンテンツの内容とステータスに記述されているハッシュ値が一致しません。",
				MimeMismatch:       "ヘッダのmimeとステータスに記述されているmimeが一致しません。",
			},
		},
	}
	Msg["en-US"] = Strings{
		Lang: "english",
		Console: ConsoleStrings{
			Starting:      "Starting Octane API Server...",
			SnowFlakeNode: "Created SnowFlake node of this process.",
			HandlerExit:   "Handler exited.",
			Error: ConsoleErrorStrings{
				DBOpen:              "Error occurred while connecting to DB.",
				DBNotSupported:      "That kind of DB is not supported.",
				SnowFlake:           "Error occurred while creating snowflake node.",
				DeadCode:            "Reached to dead code, should not be called.",
				DuplicatedID:        "Duplicated ID detected. You should make sure not to duplicate your node ids.",
				FiberResponse:       "Error occurred while fiber is responding to the request.",
				StorageOpen:         "Error occurred while opening storage.",
				StorageNotSupported: "That kind of storage is not supported.",
			},
		},
		API: APIStrings{
			APIErrorStrings{
				EmptyRoomName:      "Room name is empty.",
				EmptyDeviceName:    "Device name is empty.",
				DupDevice:          "Device is duplicated.",
				DeviceNotConnected: "Specified device is not connected.",
				BadHash:            "Hash syntax is wrong.",
				ContentMismatch:    "Mismatch detected between content and status hash.",
				MimeMismatch:       "Mismatch detected between header mime and status mime.",
			},
		},
	}
}
