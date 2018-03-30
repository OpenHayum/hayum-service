package repository

// Repositorer defines all default repository methods
type Repositorer interface {
	// SetSession sets session of the current repository
	SetSession(session interface{})

	// GetSession gets the session
	GetSession() interface{}

	// Delete deletes a model from db
	Delete(model interface{})

	// Save creates or updates a model to db
	Save(model interface{}) error

	// GetByID fetch a model by model's ID
	GetByID(id string) (interface{}, error)
}
