package kitt

type Module interface {
	Boot()
	Id() string
	Runnables() []Runnable
}
