package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/snowie2000/livetv/global"
)

func M3UGenerate() (string, error) {
	baseUrl, err := global.GetConfig("base_url")
	if err != nil {
		log.Println(err)
		return "", err
	}
	channels, err := GetAllChannel()
	if err != nil {
		log.Println(err)
		return "", err
	}
	var m3u strings.Builder
	m3u.WriteString("#EXTM3U\n")
	for _, v := range channels {
		logo := ""
		category := "LiveTV"
		if v.Category != "" {
			category = v.Category
		}
		if info, ok := global.URLCache.Load(v.URL); ok {
			logo = info.Logo
		}
		liveData := fmt.Sprintf("#EXTINF:-1, tvg-name=%s tvg-logo=%s group-title=%s, %s\n", strconv.Quote(v.Name), strconv.Quote(logo), strconv.Quote(category), v.Name)
		m3u.WriteString(liveData)
		m3u.WriteString(fmt.Sprintf("%s/live.m3u8?token=%s&c=%d\n", baseUrl, v.Token, v.ID))
	}
	return m3u.String(), nil
}
