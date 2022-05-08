package main



type huoshanFanyiRequest struct {
	Text string `json:"text"`
	Language string `json:"language"`
}


type huoshanFanyiResponse struct {
	Words []struct {
		Source int `json:"source"`
		Text string `json:"text"`
		PosList []struct {
			Type int `json:"type"`
			Phonetics []struct {
				Type int `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text string `json:"text"`
				Examples []struct {
					Type int `json:"type"`
					Sentences []struct {
						Text string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode int `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}