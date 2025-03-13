package notifications

type Notifier interface {
	Notify(notification Notification) error
}
