package config

type AppSetting struct {
	AppId         string `json:"appId" yaml:"appId"`
	AppSecret     string `json:"appSecret" yaml:"appSecret"`
	MessageToken  string `json:"messageToken" yaml:"messageToken"`
	MessageAESKey string `json:"messageAESKey" yaml:"messageAESKey"`
}

type WxmpSetting struct {
	Enabled       bool         `json:"enabled" yaml:"enabled"`
	AppId         string       `json:"appId" yaml:"appId"`
	AppSecret     string       `json:"appSecret" yaml:"appSecret"`
	MessageToken  string       `json:"messageToken" yaml:"messageToken"`
	MessageAESKey string       `json:"messageAESKey" yaml:"messageAESKey"`
	Apps          []AppSetting `json:"apps" yaml:"apps"`
}

var Setting *WxmpSetting = &WxmpSetting{
	Enabled: false,
}