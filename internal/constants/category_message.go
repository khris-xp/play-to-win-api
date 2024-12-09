package constants

const (
	CategoryCreatedSuccess     = "Category created successfully"
	CategoryUpdatedSuccess     = "Category updated successfully"
	CategoryDeletedSuccess     = "Category deleted successfully"
	CategoryRetrievedSuccess   = "Category retrieved successfully"
	CategoriesRetrievedSuccess = "Categories retrieved successfully"

	CategoryNotFoundError    = "Category not found"
	CategoryCreateError      = "Failed to create category"
	CategoryUpdateError      = "Failed to update category"
	CategoryDeleteError      = "Failed to delete category"
	CategoryInvalidIDError   = "Invalid category ID"
	CategoryInvalidDataError = "Invalid category data"
	CategoryDuplicateError   = "Category already exists"
	DatabaseError            = "Database operation failed"
	InvalidRequestError      = "Invalid request payload"
	InternalServerError      = "Internal server error"
)

const (
	StatusSuccess       = 200
	StatusCreated       = 201
	StatusBadRequest    = 400
	StatusNotFound      = 404
	StatusInternalError = 500
)
