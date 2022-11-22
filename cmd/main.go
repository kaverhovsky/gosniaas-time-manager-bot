package main

import "github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"

func main() {
	startUpLogger := common.NewLogger("development", "info")
	config := ReadConfig(".env")

}
