package comment

import (
	"github.com/Tapan176/go-practice/middleware"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func GetAllCommentsValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.BlogIDParam{},
	}
}

func AddCommentValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.BlogIDParam{},
		Body:   &requestSchema.AddCommentRequest{},
	}
}

func EditCommentValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.CommentIDParam{},
		Body:   &requestSchema.EditCommentRequest{},
	}
}

func DeleteCommentValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.CommentIDParam{},
	}
}
