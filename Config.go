package main

//Config config
type Config struct {
	Host        string   `json:"host"`
	Token       string   `json:"token"`
	Fail2banLog string   `json:"f2bLogFile"`
	Services    []string `json:"services"`
}
