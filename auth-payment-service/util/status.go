package util

type Status struct {
	Message string
}

func SendMessage(message string) Status {
	status := Status{
		Message: message,
	}
	return status
}
