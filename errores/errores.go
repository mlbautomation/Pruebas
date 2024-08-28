package errores

type message1 struct {
	message string
}

func NewMessage1(s string) message1 {
	return message1{message: s}
}

func (m message1) Error() string {
	return m.message
}

type message2 struct {
	message string
}

func NewMessage2(s string) message2 {
	return message2{message: s}
}

func (m message2) Error() (string, string) {
	return m.message, "esto no debería ser un error"
}

func VerificacionError(e error) string {
	return ("el agumento es del tipo error")
}
