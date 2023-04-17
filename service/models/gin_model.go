package models

// JSONResult ...
type JSONResult struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors []Errors    `json:"errors"`
}

// Extensions ...
type Extensions struct {
	Code       string `json:"code"`
	Field      string `json:"field"`
	StatusCode int    `json:"status_code"`
	BackToHome bool   `json:"back_to_home"`
}

// Errors ...
type Errors struct {
	Extensions Extensions `json:"extensions"`
	Message    string     `json:"message"`
}

// CustomErrors ...
type CustomErrors struct {
	Extensions       Extensions `json:"extensions"`
	Message          string     `json:"message"`
	Title            string     `json:"title,omitempty"`
	PrimaryBtnText   string     `json:"primary_btn_text,omitempty"`
	SecondaryBtnText string     `json:"secondary_btn_text,omitempty"`
	ImageUrl         string     `json:"image_url,omitempty"`
}

//CreateProductDefaultRequest ...
type CreateProductDefaultRequest struct {
	ServiceCode string `uri:"code"`
}
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data    gin.H  `json:"data,omitempty"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors,omitempty"`
}
