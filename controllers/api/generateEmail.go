package api

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"fmt"
	"strings"
	"bytes"
)

type Email struct {
	Title        string       `json:"title"`
	Context      string       `json:"context"`
}

func (as *Server) GenerateEmail(w http.ResponseWriter, r *http.Request) {
	// var emails = []Email{
	// 	{Title: "Meet the new Customer Support Representative", Context: "Dear team,\nI am pleased to introduce you to Willy who is starting today as a Customer Support Representative. She will be providing technical support and assistance to our users, making sure they enjoy the best experience with our products.\nFeel free to greet Willy in person and congratulate her with the new role!\nBest regards,\nLinda"},
	// 	{Title: "Do you have student discounts for the Annual Coding Conference?", Context: "Greetings,\nI would like to ask if you provide student discounts for tickets to the Annual Coding Conference.\nI’m a full-time student at the University of Texas and I’m very excited about your event, but unfortunately, the ticket price is too high for me. I would appreciate if you could offer me an educational discount.Looking forward to hearing from you!\nBest,\nLinda"},
	// 	{Title: "Complaint regarding the quality of the headphones", Context: "Hi there,\nI purchased the headphones at Perfect Music on Monday, August 11. Later, I discovered that the left headphone wasn’t working. Unfortunately, the staff refused to replace the headphones or return my money although I provided the receipt.\nI’m deeply disappointed about the quality of the product and the disrespectful treatment I received in your store.\nI hope to have this issue resolved and get my money back, otherwise, I will have to take further actions.Best,\nLinda"},
	// }
	// switch {
	// 	case r.Method == "GET":
	// 		JSONResponse(w, emails, http.StatusOK)
	// 	case r.Method == "POST":
			
	// 		body, _ := ioutil.ReadAll(r.Body)
	// 		keyword, _ := strconv.Unquote(string(body))
			
			
	// 		emails[0].Context =  keyword + "\n" + emails[0].Context
	// 		JSONResponse(w, emails[0], http.StatusCreated)
	// }

	switch {
		case r.Method == "GET":

			url := "https://api.nlpnchu.org/en-US/generate-phishing-email"
			
			keywords := `"golang", "website"`

			var jsonStr = []byte(`{"keywords":[`+keywords+`]}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			
			defer resp.Body.Close()

			resBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			str := string(resBody)
			result := str[2 : len(str) - 2]

			JSONResponse(w, result, http.StatusOK)

		case r.Method == "POST":
			
			body, _ := ioutil.ReadAll(r.Body)

			defer r.Body.Close()

			keyword, _ := strconv.Unquote(string(body))
			keywordList := strings.Split(keyword, " ")

			keywords := ""
			for i := 0; i < len(keywordList); i++ {
				_keyword := `"` + keywordList[i] + `"`
				if i == 0 {
					keywords += _keyword
				} else {
					keywords += `, ` + _keyword
				}
			}
			fmt.Println(keywords)

			url := "https://api.nlpnchu.org/en-US/generate-phishing-email"
			var jsonStr = []byte(`{"keywords":[`+keywords+`]}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			
			defer resp.Body.Close()

			resBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			str := string(resBody)
			result := str[2 : len(str) - 2]

			JSONResponse(w, result, http.StatusCreated)
	}
}
