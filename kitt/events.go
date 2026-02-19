package kitt

var handlers map[string][]EventHandler

type Event struct {
	Key string
	Ctx interface{}
}

type EventHandler func(e Event)

func Subscribe(key string, handler EventHandler) {
	if handlers == nil {
		handlers = make(map[string][]EventHandler)
	}
	handlers[key] = append(handlers[key], handler)
}

func Publish(e Event) {
	if handlers, exists := handlers[e.Key]; exists {
		for _, h := range handlers {
			h(e)
		}
	}
}

func GetEventContext[T interface{}](e Event) T {
	raw := e.Ctx
	ctx, ok := raw.(T)
	if ok {
		return ctx
	}

	var z T
	return z
}
