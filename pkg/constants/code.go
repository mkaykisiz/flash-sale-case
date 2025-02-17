package constants

const (
	SUCCESS       = 200
	ERROR         = 500
	INVALIDPARAMS = 400

	// auth
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10002
	ERROR_AUTH                     = 10003
	ERROR_AUTH_TOKEN               = 10004

	// product
	ERROR_ADD_PRODUCT_FAIL  = 20001
	ERROR_EDIT_PRODUCT_FAIL = 20002

	// flash sale
	ERROR_ADD_FLASH_SALE_FAIL    = 30001
	ERROR_GET_FLASH_SALE_FAIL    = 30002
	ERROR_GET_FLASH_SALES_FAIL   = 30003
	ERROR_DELETE_FLASH_SALE_FAIL = 30004
	ERROR_EDIT_FLASH_SALE_FAIL   = 30005

	//order
	ERROR_INSUFFICIENT_STOCK = 40001
	ERROR_GET_BUY_FAIL       = 40002

	//general
	INTERNAL_ERROR = 50001
)

const (
	CACHE_FLASH_SALE = "FLASH_SALE"
)
