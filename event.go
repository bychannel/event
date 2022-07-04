package event

// IEvent defined the event interface
type IEvent interface {
	Name() string
	Get(key string) interface{}
	Set(key string, val interface{})
	Add(key string, val interface{})
	Data() map[string]interface{}
	SetData(M) IEvent
	Abort(bool)
	IsAborted() bool
}

// BasicEvent define a basic event struct
type BasicEvent struct {
	name string
	data map[string]interface{}
	// mark is aborted
	aborted bool
}

// SetName set event name
func (e *BasicEvent) SetName(name string) *BasicEvent {
	e.name = name
	return e
}

// Name get event name
func (e *BasicEvent) Name() string {
	return e.name
}

// Get data by index
func (e *BasicEvent) Get(key string) interface{} {
	if v, ok := e.data[key]; ok {
		return v
	}

	return nil
}

// Set value by key
func (e *BasicEvent) Set(key string, val interface{}) {
	if e.data == nil {
		e.data = make(map[string]interface{})
	}

	e.data[key] = val
}

// Add value by key
func (e *BasicEvent) Add(key string, val interface{}) {
	if _, ok := e.data[key]; !ok {
		e.Set(key, val)
	}
}

// Data get all data
func (e *BasicEvent) Data() map[string]interface{} {
	return e.data
}

// SetData set data to the event
func (e *BasicEvent) SetData(data M) IEvent {
	if data != nil {
		e.data = data
	}
	return e
}

// Abort event loop exec
func (e *BasicEvent) Abort(abort bool) {
	e.aborted = abort
}

// IsAborted check.
func (e *BasicEvent) IsAborted() bool {
	return e.aborted
}
