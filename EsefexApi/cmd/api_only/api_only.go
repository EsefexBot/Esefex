package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/mockplayer"
	"esefexapi/sounddb/apimockdb"
	"esefexapi/util"
)

func main() {
	mockDb := apimockdb.NewApiMockDB()
	mockPlr := mockplayer.NewMockPlayer()

	api := api.NewHttpApi(mockDb, mockPlr, 8080, "esefexbot")

	<-api.Start()

	<-util.OsInterrupt()

	<-api.Stop()
}
