package blog

import (
	"net/http"
	"strconv"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
	"github.com/Tapan176/go-practice/middleware"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func GetPostsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))

	posts, err := GetPostsService(dbClient, limit, page)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, posts)
}

func CreatePostHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.CreatePostRequest)

	userID := middleware.GetUserID(request)

	post, err := CreatePostService(dbClient, req.Title, req.Body, userID, req.Category)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendCreatedResponse(response, post)
}

func UpdatePostHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	postID := request.PathValue("blogId")
	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.UpdatePostRequest)

	post, err := UpdatePostService(dbClient, postID, req.Title, req.Body, req.Category)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, post)
}

func DeletePostHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	postID := request.PathValue("blogId")

	err := DeletePostService(dbClient, postID)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendNoContentResponse(response)
}

func SearchArticleHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.SearchArticleRequest)

	limit := req.Limit
	page := req.Page
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	posts, err := SearchArticleService(dbClient, req.SearchString, limit, page)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, posts)
}

func GetArticleByIDHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	blogID := request.PathValue("blogId")

	post, err := GetArticleByIDService(dbClient, blogID)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, post)
}

func GetUserArticlesHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	userID := middleware.GetUserID(request)

	posts, err := GetUserArticlesService(dbClient, userID)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, posts)
}

func LikeArticleHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	blogID := request.PathValue("blogId")
	userID := middleware.GetUserID(request)

	result, err := LikeArticleService(dbClient, blogID, userID)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, result)
}

func DislikeArticleHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	blogID := request.PathValue("blogId")
	userID := middleware.GetUserID(request)

	result, err := DislikeArticleService(dbClient, blogID, userID)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, result)
}

func RateAuthorHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	blogID := request.PathValue("blogId")
	userID := middleware.GetUserID(request)

	// Request body is already validated by middleware
	req := request.Context().Value(constants.ValidatedBodyKey).(*struct {
		Rating int `json:"rating" validate:"required,min=1,max=5"`
	})

	result, err := RateAuthorService(dbClient, blogID, userID, req.Rating)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, result)
}

func GetArticlesByCategoryHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	categoryID := request.PathValue("categoryId")
	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	posts, err := GetArticlesByCategoryService(dbClient, categoryID, limit, page)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, posts)
}
