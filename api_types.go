package govee

type scanRespose struct {
	Msg struct {
		Cmd  string `json:"cmd"`
		Data struct {
			IP              string `json:"ip"`
			Device          string `json:"device"`
			Sku             string `json:"sku"`
			BleVersionHard  string `json:"bleVersionHard"`
			BleVersionSoft  string `json:"bleVersionSoft"`
			WifiVersionHard string `json:"wifiVersionHard"`
			WifiVersionSoft string `json:"wifiVersionSoft"`
		} `json:"data"`
	} `json:"msg"`
}
