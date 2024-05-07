package custom

import "github.com/gin-gonic/gin"

var (
	ErrDBNotConnected     = "OOPS !!!! sorry to say but you found some issue while data base connection "
	ErrDBNotPinnged       = "not able to connect the ping the connected database "
	ErrServerNOtConnected = "something went wrong, not able to reach the defined server or some other app is running on this "

	// user
	ErrUserIDNotValid           = "user id is not valid "
	ErrEmptyUserId              = "user id can't be empty"
	ErrUserNotFound             = "user not found"
	ErrUserExits                = "user already exists with this username "
	ErrUserEmailExsits          = "user already exists with this email"
	ErrUserPhoneNoExists        = "user already exists with this phone number"
	ErrSomethingWentWrong       = "OOPS !!! something went wrong, please try again sometime later"
	ErrUpdateUsernameExits      = "can't update user already exists with this username "
	ErrUpdateEmailExits         = "can't update user already exists with this Email "
	ErrUpdatePhoneNoeExits      = "can't update user already exists with this Phone number "
	ErrCanNotUpdateUser         = "could not able to update user"
	ErrUsernamePasswordMismatch = "username or password mismatch !!! try again"

	//Jwt Errors
	ErrTokenExpire        = "Auth token is expired."
	ErrInValidToken       = "Aurth token is Invalid."
	ErrUnAuthorisedAccess = "unauthorized to access this url"
	ErrTokenEmpty         = "Auth token can not be Empty , please pass valid Auth token"
	ErrInvalidAuthReq     = "Invalid authenication request"
	ErrUpdateAuthToken    = "could not able to update use auth token"

	// common
	CErrListParsing      = "error while parsing the list of data "
	CErrFindAllData      = "error occured while fetching paring list"
	CInvalidId           = "please provide a valid id."
	CInvalidParentId     = "please provide a valid parent colelction ID."
	CIdNotFoud           = "No data Found with this id."
	CParentDataNoFound   = "No parent data Found with provided parent ID."
	CInValidBody         = "please proive a valid request body"
	CDataNotSaves        = "error occured while saving data of "
	CMissingReqField     = "Request Body Incomnplete,please provoide the all the required fields"
	CInvalidStartEndDate = "please provide valid date, Invalid start or end date "
	CNotUpdated          = " not updated sucessfully"
)

func GetErrorVal(errorVal string) gin.H {
	return gin.H{"error": errorVal}
}
