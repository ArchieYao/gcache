/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package gcache

import (
	"gcache/cmd"
	gcachehttp "gcache/http"
)

func main() {
	cmd.Execute()
	gcachehttp.GenGcacheOrLoad(gcachehttp.DefaultGroupName)
	gcachehttp.StartServer()
}
