package events

var instance *Events

type Event struct {
	Message string
	X       int
	Y       int
}

type Events struct {
	infos []Event
}

func Init() {
	instance = &Events{}
}

func GetInstance() *Events {
	return instance
}

func (e *Events) CreateEvent(info string, x, y int) {
	if len(e.infos) > 50 {
		_, e.infos = e.infos[0], e.infos[1:]
	}
	e.infos = append(e.infos, Event{info, x, y})
}

func (e *Events) GetEvents() []Event {
	return e.infos
}
