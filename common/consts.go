package common

const (
	TokenStatusEnable = 1 //don't use 0, 0 is the default value!
)

const (
	ChannelStatusUnknown  = 0
	ChannelStatusEnabled  = 1 // don't use 0, 0 is the default value!
	ChannelStatusDisabled = 2 // also don't use 0
)

var UsingSQLite = false

var ChannelBaseURLs = []string{
	"",                            // 0
	"https://api.siliconflow.cn",  // 1
	"https://openai.api2d.net",    // 2
	"",                            // 3
	"https://api.openai-asia.com", // 4
	"https://api.openai-sb.com",   // 5
	"https://api.openaimax.com",   // 6
	"https://api.ohmygpt.com",     // 7
}
