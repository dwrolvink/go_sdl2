package config

// You can set config here that can be passed around in a struct
// To add a value, add an entry in the struct below, and then
// set it in GetConfig()

// To use the config some place, first import this subpackage:
//		import "go_sdl2/config"
//
// Then, you can get the struct by using:
//		var cfg = config.GetConfig()
//
// And finally, use the data:
// 		fmt.Println(cfg.ScreenTitle)

type Config struct {
	ScreenTitle string
}

func GetConfig() Config {
	return Config{
		ScreenTitle: "SDL Test Application",
	}

}