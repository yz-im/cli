package cli

type Option func(*App)

func NewNameOption(name string) Option {
	return func(a *App) {
		a.name = name
	}
}
