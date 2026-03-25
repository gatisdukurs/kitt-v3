package kernel

type App interface {
	Boot(k Kernel)
	BootController(c AppController)
	Id() string
	Runnables() []Runnable
	Kernel() Kernel
}
