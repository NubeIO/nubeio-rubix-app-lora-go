package setup

import (
	"flag"
	"path"
)

type SerialSetting struct {
	SerialPort string
	BaudRate   int
}

type AppSetting struct {
	Port      int
	GlobalDir string
	DataDir   string
	ConfigDir string
	Prod      bool
	Logging   bool
}

func App() *AppSetting {
	port := flag.Int("p", 1920, "Port")
	globalDir := flag.String("g", "", "Global Directory")
	dataDir := flag.String("d", "", "Data Directory")
	configDir := flag.String("c", "", "Config Directory")
	prod := flag.Bool("prod", false, "Deployment Mode")
	logging := flag.Bool("logging", false, "Logging")
	flag.Parse()

	appSetting := AppSetting{
		Port:      *port,
		GlobalDir: *globalDir,
		DataDir:   *dataDir,
		ConfigDir: *configDir,
		Prod:      *prod,
		Logging:   *logging,
	}
	return &appSetting
}

func Serial() *SerialSetting {
	serialSetting := SerialSetting{
		SerialPort: "/dev/ttyACM0",
		BaudRate:   38400,
	}
	return &serialSetting
}

func (appSetting *AppSetting) getAbsDataDir() string {
	if appSetting.Prod {
		return path.Join(appSetting.GlobalDir, appSetting.DataDir)
	} else {
		return path.Join("./out")
	}
}
