package websockets

type Message struct {
	Action string `json:"action"`
	Client string `json:"client"`
	Value  string `json:"value"`
}

type WsCommand struct {
	Command string `json:"cmd"`
	Code    int    `json:"c"`
	Data    *any   `json:"d"`
	Msg     string `json:"m"`
}
