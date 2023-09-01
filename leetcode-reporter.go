package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"
)

func Report(submission_ids []int, author_msg, plagiators_msg string) {

	postURL := "https://leetcode.com/contest/api/reports/"
	referer := "https://leetcode.com/contest/" + LeetcodeConfig.ContestID + "/ranking"
	client := NewLeetcodeHttpClient(LeetcodeConfig.CSRFToken, LeetcodeConfig.JWTToken)
	for _, id := range submission_ids {
		payload := []byte(fmt.Sprintf(`{
			"description": "%s",
			"submission": %d
		}`, plagiators_msg, id))

		resp, err := client.Post(postURL, referer, "application/json", bytes.NewBuffer(payload))

		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		responseBody := string(bodyBytes)

		fmt.Println("Response Body:", responseBody)

		time.Sleep(time.Millisecond * time.Duration(LeetcodeConfig.Delay))
	}

}
