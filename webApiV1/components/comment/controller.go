package comment

import (
	"net/http"
	"strconv"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
	"github.com/Tapan176/go-practice/middleware"
	"github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func GetAllCommentsController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	articleID := request.PathValue("articleId")
	postId, err := strconv.Atoi(articleID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}
	comments, err := GetAllCommentsService(dbClient, postId)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, comments)
}

func CreateCommentController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	articleID := request.PathValue("articleId")
	postId, err := strconv.Atoi(articleID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}

	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.CreateCommentRequest)
	comment, err := AddCommentService(dbClient, postId, req.UserID, req.Comment)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendCreatedResponse(response, comment)
}

func UpdateCommentController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	commentID := request.PathValue("commentId")
	commentId, err := strconv.Atoi(commentID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}
	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.UpdateCommentRequest)
	err = UpdateCommentService(dbClient, commentId, req.Comment)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendNoContentResponse(response)
}

func DeleteCommentController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	commentID := request.PathValue("commentId")
	commentId, err := strconv.Atoi(commentID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}
	err = DeleteCommentService(dbClient, commentId)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendNoContentResponse(response)
}
