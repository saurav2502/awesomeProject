package service

import (
	customLog "awesomeProject/log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {
	log := customLog.DebugLogger
	resp := "App is up"
	log.Printf("Request %v and the response is -> %v", customLog.GetRequestURI(request), resp)
	fmt.Fprintf(writer, "App is up")
}
func HandleEndPoints(writer http.ResponseWriter, req *http.Request) {
	infolog := customLog.InfoLogger
	debugLog := customLog.DebugLogger
	infolog.Print("entering into HandleEndPoints")
	var appInfos = []*Appinfo{}
	appInfo := new(Appinfo)
	appInfo.AppName = "Spotify"
	appInfo.AppId = "spotify-inst"
	appInfo.AppFileName = "spotify_en.txt"
	appInfos = append(appInfos, appInfo)
	appInfo.AppName = "geo"
	appInfo.AppId = "geo-inst"
	appInfo.AppFileName = "geo_en.txt"
	appInfos = append(appInfos, appInfo)
	appInfo.AppName = "airtel"
	appInfo.AppId = "airtel-inst"
	appInfo.AppFileName = "airtel_en.txt"
	appInfos = append(appInfos, appInfo)
	//json.NewEncoder(writer).Encode(appInfos)
	infolog.Printf("Returning from info log with res: %v", PrettyString(appInfos))
	debugLog.Printf("Request with URI : %v , and response : %v", customLog.GetRequestURI(req), PrettyString(appInfos))
	json.NewEncoder(writer).Encode(appInfos)
}

type Appinfo struct {
	AppName     string `json:"app_name"`
	AppId       string `json:"app_id"`
	AppFileName string `json:"app_file_name"`
}

func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func PrettyString(datas interface{}) string {
	b, _ := json.Marshal(datas)
	return string(b)
}
