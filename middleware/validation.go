package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ValidationSchema struct {
	Body   interface{}
	Query  interface{}
	Params interface{}
}

func ValidationMiddleware(schema ValidationSchema) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validationErrors := make(map[string]string)
			ctx := r.Context()

			if schema.Body != nil {
				schemaType := reflect.TypeOf(schema.Body)
				if schemaType.Kind() == reflect.Ptr {
					schemaType = schemaType.Elem()
				}
				instance := reflect.New(schemaType).Interface()

				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					internal.HandleError(w, "invalid_request_body")
					return
				}

				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				if err := json.Unmarshal(bodyBytes, instance); err != nil {
					internal.HandleError(w, "invalid_request_body")
					return
				}

				if err := validate.Struct(instance); err != nil {
					if validatorErrs, ok := err.(validator.ValidationErrors); ok {
						for _, e := range validatorErrs {
							validationErrors["body."+e.Field()] = formatValidationError(e)
						}
					}
				}

				ctx = context.WithValue(ctx, constants.ValidatedBodyKey, instance)
			}

			if schema.Query != nil {
				schemaType := reflect.TypeOf(schema.Query)
				if schemaType.Kind() == reflect.Ptr {
					schemaType = schemaType.Elem()
				}
				instance := reflect.New(schemaType).Elem()

				queryValues := r.URL.Query()
				for i := 0; i < schemaType.NumField(); i++ {
					field := schemaType.Field(i)
					jsonTag := field.Tag.Get("json")
					if jsonTag != "" && jsonTag != "-" {
						if value := queryValues.Get(jsonTag); value != "" {
							instance.Field(i).SetString(value)
						}
					}
				}

				if err := validate.Struct(instance.Interface()); err != nil {
					if validatorErrs, ok := err.(validator.ValidationErrors); ok {
						for _, e := range validatorErrs {
							validationErrors["query."+e.Field()] = formatValidationError(e)
						}
					}
				}
			}

			if schema.Params != nil {
				schemaType := reflect.TypeOf(schema.Params)
				if schemaType.Kind() == reflect.Ptr {
					schemaType = schemaType.Elem()
				}
				instance := reflect.New(schemaType).Elem()

				for i := 0; i < schemaType.NumField(); i++ {
					field := schemaType.Field(i)
					jsonTag := field.Tag.Get("json")
					if jsonTag != "" && jsonTag != "-" {
						if value := r.PathValue(jsonTag); value != "" {
							instance.Field(i).SetString(value)
						}
					}
				}

				if err := validate.Struct(instance.Interface()); err != nil {
					if validatorErrs, ok := err.(validator.ValidationErrors); ok {
						for _, e := range validatorErrs {
							validationErrors["params."+e.Field()] = formatValidationError(e)
						}
					}
				}
			}

			if len(validationErrors) > 0 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"code":    "validation_failed",
					"message": "Validation failed",
					"details": validationErrors,
				})
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func formatValidationError(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "min":
		return field + " must be at least " + e.Param() + " characters"
	case "max":
		return field + " must be at most " + e.Param() + " characters"
	case "eqfield":
		return field + " must match " + e.Param()
	case "numeric":
		return field + " must be numeric"
	case "alpha":
		return field + " must contain only letters"
	case "alphanum":
		return field + " must contain only letters and numbers"
	default:
		return field + " is invalid"
	}
}
