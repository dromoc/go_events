package message_queue

// EventListener describes an interface for a class that can listen to events.
type EventListener interface {
	Listen(events ...string) (<-chan Event, <-chan error, error)
}