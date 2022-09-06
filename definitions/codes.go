package definitions

const (
	//Error
	ERR_DUP_DEVICE             = "ERR_DUP_DEVICE"             //Connected with duplicated device name
	ERR_BAD_REQUEST            = "ERR_BAD_REQUEST"            // Generic bad request
	ERR_INVALID_ID             = "ERR_INVALID_ID"             // Room ID is invalid
	ERR_INTERNAL_EXCEPTION     = "ERR_INTERNAL_EXCEPTION"     // Server side error
	ERR_NO_SUCH_ROOM           = "ERR_NO_SUCH_ROOM"           // No such room exists
	ERR_NO_STATUS_SET          = "ERR_NO_STATUS_SET"          // Status is not set
	ERR_DEVICE_NOT_CONNECTED   = "ERR_DEVICE_NOT_CONNECTED"   // Device is not connected
	ERR_CONTENT_MISMATCH       = "ERR_CONTENT_MISMATCH"       // Content crc hash or mime is wrong
	ERR_BAD_HASH               = "ERR_BAD_HASH"               // Hash style is wrong
	ERR_UNAUTHORIZED           = "ERR_UNAUTHORIZED"           // Bad token
	ERR_BAD_MIME               = "ERR_BAD_MIME"               //Mime type is not accepted
	ERR_BAD_CONNECTION_REQUEST = "ERR_BAD_CONNECTION_REQUEST" // Invalid connection request
	//Enum
	Faulty    = "faulty"
	Degraded  = "degraded"
	Healthy   = "healthy"
	Clipboard = "clipboard"
	File      = "file"
	MultiFile = "multi-file"
	//Headers
	HeaderToken = "X-Octane-API-Token"
	//Connect requests
	RequestConnect    = "connect"
	RequestDisconnect = "disconnect"
)

var TypeClipboard []string
var TypeMultiFile []string

func init() {
	TypeClipboard = []string{"text/plain;charset=utf-8", "text/plain"}
	TypeMultiFile = []string{"application/x-lzma", "application/x-7z-compressed"}
}
