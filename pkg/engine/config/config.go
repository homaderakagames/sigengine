package config

type Settings struct {
	ScreenWidth  int `env:"SETTINGS_SCREEN_WIDTH" envDefault:"1280"`
	ScreenHeight int `env:"SETTINGS_SCREEN_HEIGHT" envDefaul:"720"`
}

type Common struct {
	Name    string `env:"GAME_NAME" envDefault:"Undefined"`
	Version string `env:"VERSION" envDefault:"Undefined"`
}
