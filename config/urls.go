package config

func GetUrl(routes string) string {
	return baseUrl + routes + suffix
}
