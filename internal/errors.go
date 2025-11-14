package internal

import "net/http"

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	PasswordAndConfirmPasswordDoNotMatch = &AppError{
		Code:       "password_and_confirm_password_do_not_match",
		Message:    "Password and confirm password do not match",
		StatusCode: http.StatusBadRequest,
	}
	UserNotFound = &AppError{
		Code:       "user_not_found",
		Message:    "User not found",
		StatusCode: http.StatusNotFound,
	}
	IncorrectPassword = &AppError{
		Code:       "incorrect_password",
		Message:    "Incorrect password",
		StatusCode: http.StatusUnauthorized,
	}
	AuthenticationFailed = &AppError{
		Code:       "authentication_failed",
		Message:    "Authentication failed",
		StatusCode: http.StatusInternalServerError,
	}
	UserHasNotLogin = &AppError{
		Code:       "user_has_not_login",
		Message:    "User has not login",
		StatusCode: http.StatusBadRequest,
	}
	LogoutFailed = &AppError{
		Code:       "logout_failed",
		Message:    "Logout failed",
		StatusCode: http.StatusInternalServerError,
	}
	UserAlreadyExists = &AppError{
		Code:       "user_already_exists",
		Message:    "User already exists",
		StatusCode: http.StatusConflict,
	}
	FailedToRegisterUser = &AppError{
		Code:       "failed_to_register_user",
		Message:    "Failed to register user",
		StatusCode: http.StatusInternalServerError,
	}
	InvalidToken = &AppError{
		Code:       "invalid_token",
		Message:    "Invalid token",
		StatusCode: http.StatusUnauthorized,
	}
	InvalidVerification = &AppError{
		Code:       "invalid_verification",
		Message:    "Invalid verification",
		StatusCode: http.StatusUnauthorized,
	}
	NewEmailMustBeDifferentFromCurrentEmail = &AppError{
		Code:       "new_email_must_be_different_from_current_email",
		Message:    "New email must be different from current email",
		StatusCode: http.StatusBadRequest,
	}
	EmailAlreadyUsed = &AppError{
		Code:       "email_already_used",
		Message:    "Email already used by another user",
		StatusCode: http.StatusConflict,
	}
	NewPasswordSameAsOldPassword = &AppError{
		Code:       "new_password_same_as_old_password",
		Message:    "New password cannot be the same as the old password",
		StatusCode: http.StatusBadRequest,
	}
	FailedToSendEmail = &AppError{
		Code:       "failed_to_send_email",
		Message:    "Failed to send email",
		StatusCode: http.StatusInternalServerError,
	}

	AddBlogFailed = &AppError{
		Code:       "add_blog_failed",
		Message:    "Failed to add blog",
		StatusCode: http.StatusInternalServerError,
	}
	FetchBlogsFailed = &AppError{
		Code:       "fetch_blogs_failed",
		Message:    "Failed to fetch blogs",
		StatusCode: http.StatusInternalServerError,
	}
	UpdateBlogFailed = &AppError{
		Code:       "update_blog_failed",
		Message:    "Failed to update blog",
		StatusCode: http.StatusInternalServerError,
	}
	DeleteBlogFailed = &AppError{
		Code:       "delete_blog_failed",
		Message:    "Failed to delete blog",
		StatusCode: http.StatusInternalServerError,
	}
	BlogAlreadyExist = &AppError{
		Code:       "blog_already_exist",
		Message:    "Blog already exists for the user",
		StatusCode: http.StatusConflict,
	}
	BlogNotFound = &AppError{
		Code:       "blog_not_found",
		Message:    "Blog not found",
		StatusCode: http.StatusNotFound,
	}
	BlogNotAuthorized = &AppError{
		Code:       "blog_not_authorized",
		Message:    "You are not authorized to update or delete this blog",
		StatusCode: http.StatusForbidden,
	}

	CommentNotFound = &AppError{
		Code:       "comment_not_found",
		Message:    "Comment not found",
		StatusCode: http.StatusNotFound,
	}
	FetchCommentFailed = &AppError{
		Code:       "fetch_comment_failed",
		Message:    "Failed to fetch comment",
		StatusCode: http.StatusInternalServerError,
	}
	IncreaseQuantityFailed = &AppError{
		Code:       "increase_quantity_failed",
		Message:    "Failed to increase quantity",
		StatusCode: http.StatusInternalServerError,
	}
	DecreaseQuantityFailed = &AppError{
		Code:       "decrease_quantity_failed",
		Message:    "Failed to decrease quantity",
		StatusCode: http.StatusInternalServerError,
	}
	DeleteBlogFromCommentFailed = &AppError{
		Code:       "delete_blog_from_comment_failed",
		Message:    "Failed to delete blog from comment",
		StatusCode: http.StatusInternalServerError,
	}

	InvalidCategoryID = &AppError{
		Code:       "invalid_category_id",
		Message:    "Invalid category ID",
		StatusCode: http.StatusBadRequest,
	}
	CategoryNotFound = &AppError{
		Code:       "category_not_found",
		Message:    "Category not found",
		StatusCode: http.StatusNotFound,
	}

	InvalidSearchString = &AppError{
		Code:       "invalid_search_string",
		Message:    "Invalid search string",
		StatusCode: http.StatusBadRequest,
	}

	Unauthorized = &AppError{
		Code:       "unauthorized",
		Message:    "You are not authorized",
		StatusCode: http.StatusUnauthorized,
	}
	Forbidden = &AppError{
		Code:       "forbidden",
		Message:    "Insufficient permissions to access this resource",
		StatusCode: http.StatusForbidden,
	}
	PleaseLogin = &AppError{
		Code:       "please_login",
		Message:    "Please login",
		StatusCode: http.StatusUnauthorized,
	}

	InvalidRequestBody = &AppError{
		Code:       "invalid_request_body",
		Message:    "Invalid request body",
		StatusCode: http.StatusBadRequest,
	}
	InternalServerError = &AppError{
		Code:       "internal_server_error",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
)

var ErrorMap = map[string]*AppError{
	"password_and_confirm_password_do_not_match": PasswordAndConfirmPasswordDoNotMatch,
	"user_not_found":                                 UserNotFound,
	"incorrect_password":                             IncorrectPassword,
	"authentication_failed":                          AuthenticationFailed,
	"user_has_not_login":                             UserHasNotLogin,
	"logout_failed":                                  LogoutFailed,
	"user_already_exists":                            UserAlreadyExists,
	"failed_to_register_user":                        FailedToRegisterUser,
	"invalid_token":                                  InvalidToken,
	"invalid_verification":                           InvalidVerification,
	"new_email_must_be_different_from_current_email": NewEmailMustBeDifferentFromCurrentEmail,
	"email_already_used":                             EmailAlreadyUsed,
	"new_password_same_as_old_password":              NewPasswordSameAsOldPassword,
	"failed_to_send_email":                           FailedToSendEmail,
	"add_blog_failed":                                AddBlogFailed,
	"fetch_blogs_failed":                             FetchBlogsFailed,
	"update_blog_failed":                             UpdateBlogFailed,
	"delete_blog_failed":                             DeleteBlogFailed,
	"blog_already_exist":                             BlogAlreadyExist,
	"blog_not_found":                                 BlogNotFound,
	"blog_not_authorized":                            BlogNotAuthorized,
	"comment_not_found":                              CommentNotFound,
	"fetch_comment_failed":                           FetchCommentFailed,
	"increase_quantity_failed":                       IncreaseQuantityFailed,
	"decrease_quantity_failed":                       DecreaseQuantityFailed,
	"delete_blog_from_comment_failed":                DeleteBlogFromCommentFailed,
	"invalid_category_id":                            InvalidCategoryID,
	"category_not_found":                             CategoryNotFound,
	"invalid_search_string":                          InvalidSearchString,
	"unauthorized":                                   Unauthorized,
	"forbidden":                                      Forbidden,
	"please_login":                                   PleaseLogin,
	"invalid_request_body":                           InvalidRequestBody,
	"internal_server_error":                          InternalServerError,
}

func GetError(code string) *AppError {
	if err, exists := ErrorMap[code]; exists {
		return err
	}
	return InternalServerError
}
