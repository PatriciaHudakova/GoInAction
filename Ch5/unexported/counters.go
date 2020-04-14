package unexported

type alertCounter int //unexported type

func New(value int) alertCounter { //factory
	return alertCounter(value)
}
