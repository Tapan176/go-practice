package blog

import (
	"github.com/Tapan176/go-practice/middleware"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func GetArticlesByIDValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: requestSchema.BlogIDParam{},
	}
}

func GetArticlesByCategoryValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: requestSchema.CategoryIDParam{},
		Query:  requestSchema.GetBlogsQuery{},
	}
}

func AddArticleValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.CreatePostRequest{},
	}
}

func EditArticleValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: requestSchema.BlogIDParam{},
		Body:   requestSchema.UpdatePostRequest{},
	}
}

func DeleteArticleValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: requestSchema.BlogIDParam{},
	}
}

func SearchArticleValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.SearchArticleRequest{},
	}
}

func GetArticlesValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Query: requestSchema.GetBlogsQuery{},
	}
}
