package utils

type Message string

const (
	InvalidFirstNameMessage Message = "Invalid first name."
	InvalidLastNameMessage  Message = "Invalid last name."
	InvalidUsernameMessage  Message = "Invalid username."
	InvalidEmailMessage     Message = "Invalid email."
	InvalidPasswordMessage  Message = "Invalid password."
	UnknownMessage          Message = "Unknown"
)

func (enum Message) Value() string {
	switch enum {
	case InvalidFirstNameMessage:
		return "Invalid first name."
	case InvalidLastNameMessage:
		return "Invalid last name."
	case InvalidUsernameMessage:
		return "Invalid username."
	case InvalidEmailMessage:
		return "Invalid email."
	case InvalidPasswordMessage:
		return "Invalid password."
	case UnknownMessage:
		return "Unknown"
	}

	return "Unknown"
}
