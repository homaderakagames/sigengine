package config

type Settings struct {
	ScreenWidth  int
	ScreenHeight int
}

type Common struct {
	Name    string ``
	Version string

	Settings Settings
}
