package zsimple

type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	// misal, kita ingin menjadikan fiel configuration ini yang ada dalam struct Application sebagai dependency
	return &Application{
		Configuration: &Configuration{
			Name: "Vin-Barbatos",
		},
	}
}
