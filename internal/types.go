package internal

type Response struct {
	Message string      `json:"message"`
	Result  []BlockData `json:"result"`
}

type BlockData struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
}

type Config struct {
	Validators []ValidatorData `json:"validators"`
	Interval   int             `json:"interval"`
}

type ValidatorData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	URL     string `json:"url"`
}

type Validator struct {
	Name        string
	MinedBlocks int
}
