package udp

func generateMap() {
	var mapCode = make(map[int]string)
	mapCode[65] = "Created"
	mapCode[66] = "Deleted"
	mapCode[67] = "Valid"
	mapCode[68] = "Changed"
	mapCode[69] = "Content"
	mapCode[128] = "BadRequest"
	mapCode[129] = "Unauthorized"
	mapCode[130] = "BadOption"
	mapCode[131] = "Forbidden"
	mapCode[132] = "NotFound"
	mapCode[133] = "MethodNotAllowed"
	mapCode[134] = "NotAcceptable"
	mapCode[140] = "PreconditionFailed"
	mapCode[141] = "RequestEntityTooLarge"
	mapCode[143] = "UnsupportedMediaType"
	mapCode[160] = "InternalServerError"
	mapCode[161] = "NotImplemented"
	mapCode[162] = "BadGateway"
	mapCode[163] = "ServiceUnavailable"
	mapCode[164] = "GatewayTimeout"
	mapCode[165] = "ProxyingNotSupported"
	mapType := make(map[int]string, 4)
	mapType[0] = "Confirmable"
	mapType[1] = "NonConfirmable"
	mapType[2] = "Acknowledgement"
	mapType[3] = "Reset"
}
