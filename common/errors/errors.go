package errors

type UserAlreadyExistsError struct{}

func (m *UserAlreadyExistsError) Error() string {
	return "A user with this username already exists"
}
