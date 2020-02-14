package main

import (
	"log"
	"os"

	dbhelper "github.com/JojiiOfficial/GoDBHelper"
	"gopkg.in/alecthomas/kingpin.v2"
)

const appName = "server"

var showTimeInLog = false
var logPrefix = ""

var (
	app        = kingpin.New(appName, "A Rest server")
	appDebug   = app.Flag("debug", "Enable debug mode").Short('d').Bool()
	appNoColor = app.Flag("no-color", "Disable colors").Envar(getEnVar(EnVarNoColor)).Bool()
	appYes     = app.Flag("yes", "Skips confirmations").Short('y').Envar(getEnVar(EnVarYes)).Bool()
	appCfgFile = app.
			Flag("config", "the configuration file for the subscriber").
			Envar(getEnVar(EnVarConfigFile)).
			Short('c').String()

	//Server commands
	//Server start
	serverCmd      = app.Command("server", "Commands for the server")
	serverCmdStart = serverCmd.Command("start", "Start the server")
)

func main() {
	app.HelpFlag.Short('h')
	app.Version("0.01")

	//parsing the args
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

	var (
		config *ConfigStruct
		db     *dbhelper.DBhelper
	)

	var shouldExit bool
	config, shouldExit = InitConfig(*appCfgFile, false)
	if shouldExit {
		return
	}

	if !config.Check() {
		if *appDebug {
			log.Println("Exiting")
		}
		return
	}

	var err error
	db, err = connectDB(config)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	switch parsed {
	//Server --------------------
	case serverCmdStart.FullCommand():
		{
			runCmd(config, db)
		}
	}
}
