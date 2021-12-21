package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

// golang-migrate/migrateをパッケージとしてGoから使ってみる
// https://qiita.com/tanden/items/7b4fb1686a61dd5f580d#golang-migratemigrate%E3%81%A8%E3%81%AF

// 【Golang】migrateでDBマイグレーションをする
// https: //simple-minds-think-alike.moritamorie.com/entry/golang-migrate#%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%83%87%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E6%A7%8B%E6%88%90%E7%A2%BA%E8%AA%8D

// sql-migrateとGormでマイグレーションとシードを実装する
// https://ryotarch.com/go-lang/implement-migrate-and-seed-with-sql-migrate-and-gorm/

//sql and database info
const (
	Source = "file://./migrations/"
	// Database = "mysql://user:password@tcp(0.0.0.0:3306)/database"
	Database = "mysql://echo:golang@tcp(mysql:3306)/echo"
)

//declare command line options
var (
	Command = flag.String("exec", "", "set up or down as a argument")
	Force   = flag.Bool("f", false, "force exec fixed sql")
)

//available command list
var AvailableExecCommands = map[string]string{
	"up":      "Execute up sqls",
	"down":    "Execute down sqls",
	"version": "Just check current migrate version",
}

func main() {
	flag.Parse()
	if len(*Command) < 1 {
		fmt.Println("\nerror: no argument\n")
		showUsageMessge()
		os.Exit(1)
		return
	}

	m, err := migrate.New(Source, Database)
	if err != nil {
		fmt.Println("err", err)
	}
	version, dirty, err := m.Version()
	showVersionInfo(version, dirty, err)

	fmt.Println("command: exec", *Command)
	applyQuery(m, version, dirty)
}

//exec up or down sqls
//with force option if needed
func applyQuery(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *Force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	var err error
	switch *Command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		//do nothing
		return
	default:
		fmt.Println("\nerror: invalid command '" + *Command + "'\n")
		showUsageMessge()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *Command+"\n")
		fmt.Println("updated version info")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func showUsageMessge() {
	fmt.Println("-------------------------------------")
	fmt.Println("Usage")
	fmt.Println("  go run migrate.go -exec <command>\n")
	fmt.Println("Available Exec Commands: ")
	for available_command, detail := range AvailableExecCommands {
		fmt.Println("  " + available_command + " : " + detail)
	}
	fmt.Println("-------------------------------------")
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version  : ", version)
	fmt.Println("dirty    : ", dirty)
	fmt.Println("error    : ", err)
	fmt.Println("-------------------")
}