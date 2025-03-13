package policies

type ModelPolicy[T any] interface {
	ViewAny() bool
	View(model T) bool
	Create() bool
	Delete(model T) bool
	Update(model T) bool
}
