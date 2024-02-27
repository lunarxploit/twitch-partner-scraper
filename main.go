package main

import (
	"atomic/imports"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gookit/color"
	"github.com/valyala/fasthttp"
)

var (
	idChan chan int
)

func main() {
	color.Println("<fg=blue>[+]</> Starting Twitch Scraper")
	valid := make(chan string)
	idChan = make(chan int)

	go func() {
		for username := range valid {
			go getLastBroadcast(username)
		}
	}()

	go incrementID()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				getUser(<-idChan, valid)
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	select {}
}

func incrementID() {
	id := 9999
	for {
		id++
		idChan <- id
	}
}

func getUser(user int, valid chan<- string) {
	client := &fasthttp.Client{}
	var data = (`[{"operationName":"ReportUserModal_ReportWizardData","variables":{"targetUserID":"` + strconv.Itoa(user) + `","reportSessionID":"e7a8c33e47f12b2df09d508d47b98b77","reportWizardVersion":"1.0"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"53e3604167556125620dfc8ceed34ab2a14fb1bdbd16a66c62a3db6a2e862ea8"}}}]`)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://gql.twitch.tv/gql")
	req.Header.SetMethod("POST")
	req.SetBodyString(data)

	req.Header.Set("client-id", "kimne78kx3ncx6brgo4mv6wki5h1ko")
	req.Header.Set("client-session-id", "f9e3bc910b382c99")
	req.Header.Set("client-version", "9cd238f2-7208-4f5c-ad92-9f3bb7cfbc43")
	req.Header.Set("content-type", "text/plain;charset=UTF-8")
	req.Header.Set("x-device-id", "tgGvMHrCEQOS7DoYx6i7b5XCiz1fcfXn")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	client.Do(req, resp)

	var userInfo imports.UserInfo
	json.Unmarshal(resp.Body(), &userInfo)

	if userInfo[0].Data.TargetUser.DisplayName != "" {
		valid <- userInfo[0].Data.TargetUser.DisplayName
	}
}

func getLastBroadcast(login string) {
	client := &fasthttp.Client{}
	var data = `[{"operationName":"StreamSchedule","variables":{"login":"` + login + `","startingWeekday":"MONDAY","utcOffsetMinutes":60,"startAt":"2024-02-25T23:00:00.000Z","endAt":"2024-03-03T22:59:59.059Z"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"83552f5614707fd3e897495c18875b6fa9c83d8cf11e73b9f158f3173b4f3b75"}}},{"operationName":"ChannelAvatar","variables":{"channelLogin":"` + login + `"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"84ed918aaa9aaf930e58ac81733f552abeef8ac26c0117746865428a7e5c8ab0"}}}]`
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://gql.twitch.tv/gql")
	req.Header.SetMethod("POST")
	req.SetBodyString(data)

	req.Header.Set("client-id", "kimne78kx3ncx6brgo4mv6wki5h1ko")
	req.Header.Set("client-session-id", "f9e3bc910b382c99")
	req.Header.Set("client-version", "9cd238f2-7208-4f5c-ad92-9f3bb7cfbc43")
	req.Header.Set("content-type", "text/plain;charset=UTF-8")
	req.Header.Set("x-device-id", "tgGvMHrCEQOS7DoYx6i7b5XCiz1fcfXn")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	client.Do(req, resp)

	var response []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		fmt.Println("[!]", err)
		return
	}

	var isPartner bool
	var lastBroadcast time.Time

	for _, item := range response {
		if data, ok := item["data"].(map[string]interface{}); ok {
			if userData, ok := data["user"].(map[string]interface{}); ok {
				if user, ok := userData["isPartner"].(bool); ok {
					isPartner = user
				}
				if broadcast, ok := userData["lastBroadcast"].(map[string]interface{}); ok {
					if startedAtStr, ok := broadcast["startedAt"].(string); ok {
						lastBroadcast, _ = time.Parse(time.RFC3339Nano, startedAtStr)
					}
				}
			}
		}
	}

	if isPartner && !lastBroadcast.IsZero() && time.Since(lastBroadcast).Hours() > float64(5*365*24) {
		color.Printf("<fg=4ade80>[+]</> %s :: <fg=616161>%v</> :: <fg=9333ea>Partnered</>\n", login, lastBroadcast.Format("2006-01-02 15:04:05"))
	} else if isPartner {
		color.Printf("<fg=f87171>[!]</> %s :: <fg=616161>%v</> :: <fg=f87171>Too New</>\n", login, lastBroadcast.Format("2006-01-02 15:04:05"))
	}

}
