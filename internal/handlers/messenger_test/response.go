package messenger_test

type Response struct {
	Success        bool          `json:"success"`
	Message        string        `json:"message"`
	AdditionalData ResponseInner `json:"addithionalData"`
}

type ResponseInner struct {
	AdditionalInfo string `json:"additionalInfo"`
}
