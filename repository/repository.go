package repository

// Repositorer defines all default repository methods
type Repositorer interface {
	// Save creates or updates a model to db
	Save(model interface{}) error

	// GetByID fetch a model by model's ID
	GetByID(id string, model interface{}) error

	// Delete deletes a model from db
	Delete(model interface{}) error

	// DeleteByID deletes a model by Id
	DeleteByID(id string) error
}
