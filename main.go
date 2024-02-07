/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"ashu_music/cmd"
	"ashu_music/conf"
	"ashu_music/src"
)

func main() {
	conf.Init()
	src.Init()
	cmd.Execute()
}
