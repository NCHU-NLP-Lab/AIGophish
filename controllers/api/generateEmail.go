package api

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"fmt"
	"strings"
	"bytes"
)

func (as *Server) GenerateEmail(w http.ResponseWriter, r *http.Request) {

	switch {
		case r.Method == "GET":

			url := "https://api.nlpnchu.org/en-US/generate-phishing-email"
			
			subject := `"Test"`
			category := `"TECH"`
			keywords := `"golang", "website"`

			var jsonStr = []byte(`{"keywords":[`+keywords+`], "category":`+category+`, "title":`+subject+`}`)
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

			data, _ := strconv.Unquote(string(body))
			
			dataList := strings.Split(data, "///")

			subject := dataList[0]
			category := dataList[1]
			keywordList := strings.Split(dataList[2], ", ")
			
			keywords := ""
			for i := 0; i < len(keywordList); i++ {
				_keyword := `"` + keywordList[i] + `"`
				if i == 0 {
					keywords += _keyword
				} else {
					keywords += `, ` + _keyword
				}
			}
			// fmt.Println(keywords)

			url := "https://api.nlpnchu.org/en-US/generate-phishing-email"
			var jsonStr = []byte(`{"keywords":[`+keywords+`], "category":"`+category+`", "title":"`+subject+`"}`)
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
