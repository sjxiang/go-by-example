package main



// 示例 {trans_type: "en2zh", source: "good"}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source string `json:"source"`
	// UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}

