package expedition

import "errors"

// Function for defining error messages.
var (
	ErrCode400 = errors.New("400 Bad Request")
	ErrCode401 = errors.New("401 Unauthorized")
	ErrCode402 = errors.New("402 Payment Required ")
	ErrCode403 = errors.New("403 Forbidden")
	ErrCode404 = errors.New("404 Not Found")
	ErrCode405 = errors.New("405 Method Not Allowed")
	ErrCode406 = errors.New("406 Not Acceptable")
	ErrCode407 = errors.New("407 Proxy Authentication Required")
	ErrCode408 = errors.New("408 Request Timeout")
	ErrCode409 = errors.New("409 Conflict")
	ErrCode410 = errors.New("410 Gone")
	ErrCode411 = errors.New("411 Length Required")
	ErrCode412 = errors.New("412 Precondition Failed")
	ErrCode413 = errors.New("413 Payload Too Large")
	ErrCode414 = errors.New("414 URI Too Long")
	ErrCode415 = errors.New("415 Unsupported Media Type")
	ErrCode416 = errors.New("416 Range Not Satisfiable")
	ErrCode417 = errors.New("417 Expectation Failed")
	ErrCode418 = errors.New("418 I'm a teapot")
	ErrCode421 = errors.New("421 Misdirected Request")
	ErrCode422 = errors.New("422 Unprocessable Entity")
	ErrCode423 = errors.New("423 Locked")
	ErrCode424 = errors.New("424 Failed Dependency")
	ErrCode425 = errors.New("425 Too Early")
	ErrCode426 = errors.New("426 Upgrade Required")
	ErrCode428 = errors.New("428 Precondition Required")
	ErrCode429 = errors.New("429 Too Many Requests")
	ErrCode431 = errors.New("431 Request Header Fields Too Large")
	ErrCode451 = errors.New("451 Unavailable For Legal Reasons")
	ErrCode500 = errors.New("500 Internal Server Error")
	ErrCode501 = errors.New("501 Not Implemented")
	ErrCode502 = errors.New("502 Bad Gateway")
	ErrCode503 = errors.New("503 Service Unavailable")
	ErrCode504 = errors.New("504 Gateway Timeout")
	ErrCode505 = errors.New("505 HTTP Version Not Supported")
	ErrCode506 = errors.New("506 Variant Also Negotiates")
	ErrCode507 = errors.New("507 Insufficient Storage")
	ErrCode508 = errors.New("508 Loop Detected")
	ErrCode510 = errors.New("510 Not Extended")
	ErrCode511 = errors.New("511 Network Authentication Required")
)

// getErrorMessage is a function to return a message according to the error code.
func getErrorMessage(errCode int) error {
	switch errCode {
	case 400:
		return ErrCode400
	case 401:
		return ErrCode401
	case 402:
		return ErrCode402
	case 403:
		return ErrCode403
	case 404:
		return ErrCode404
	case 405:
		return ErrCode405
	case 406:
		return ErrCode406
	case 407:
		return ErrCode407
	case 408:
		return ErrCode408
	case 409:
		return ErrCode409
	case 410:
		return ErrCode410
	case 411:
		return ErrCode411
	case 412:
		return ErrCode412
	case 413:
		return ErrCode413
	case 414:
		return ErrCode414
	case 415:
		return ErrCode415
	case 416:
		return ErrCode416
	case 417:
		return ErrCode417
	case 418:
		return ErrCode418
	case 421:
		return ErrCode421
	case 422:
		return ErrCode422
	case 423:
		return ErrCode423
	case 424:
		return ErrCode424
	case 425:
		return ErrCode425
	case 426:
		return ErrCode426
	case 428:
		return ErrCode428
	case 429:
		return ErrCode429
	case 431:
		return ErrCode431
	case 451:
		return ErrCode451
	case 500:
		return ErrCode500
	case 501:
		return ErrCode501
	case 502:
		return ErrCode502
	case 503:
		return ErrCode503
	case 504:
		return ErrCode504
	case 505:
		return ErrCode505
	case 506:
		return ErrCode506
	case 507:
		return ErrCode507
	case 508:
		return ErrCode508
	case 510:
		return ErrCode510
	case 511:
		return ErrCode511
	}
	return nil
}
