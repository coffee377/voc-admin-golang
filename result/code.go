package result

var (
	// Common errors
	OK                  = NewBizError("0", "OK")
	InternalServerError = NewBizError("500", "Internal server error.")
	Unauthorized        = NewBizError("401", "Unauthorized")
	NotFound            = NewBizError("404", "Not Found")

	// user errors
)
