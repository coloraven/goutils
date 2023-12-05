package goutils

import (
	"fmt"
	"testing"
)

func TestWriteMapsToXLSX(t *testing.T) {
	datas := []map[string]string{
		{"Name": "Alice", "City": `New\n,," York`, "Age": "30"},
		{"Name": "Bob", "City": "Los Angeles", "Age": "25"},
	}
	WriteMapsToXLSX(datas, "sxs")
}

func TestWriteMapsToCSV(t *testing.T) {
	datas := []map[string]string{
		{"Name": "Alice", "City": `New\n,," York`, "Age": "30"},
		{"Name": "Bob", "City": "Los Angeles", "Age": "25"},
	}
	WriteMapsToCSV(datas, "sxs")
}

func TestStringHeadersToMap(t *testing.T) {
	httpHeaders := `	Host: github.githubassets.com
									User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0
									Accept: */*
									Accept-Language: zh-CN,en-US;q=0.7,en;q=0.3
									Accept-Encoding: gzip, deflate, br
									Origin: https://github.com
									Connection: keep-alive
									Referer: https://github.com/
									Sec-Fetch-Dest: script
									Sec-Fetch-Mode: cors
									Sec-Fetch-Site: cross-site
									Pragma: no-cache
									Cache-Control: no-cache`
	m, _ := StringHeadersToMap(httpHeaders)
	for k, v := range m {
		fmt.Printf("%s:%s\n", k, v)
	}
}

func TestStringCookiesToMAP(t *testing.T) {
	cookies := `_octo=GH1.1.1468507307.1691830472; logged_in=yes; _device_id=53992bd74bc36a685d2c0af1ca953353; user_session=DiJlNmBvedNA5SiNGoUdapfN1Vg3ebYX6Y38Q-CoNtAONSy9; dotcom_user=coloraven; has_recent_activity=1; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22light%22%2C%22color_mode%22%3A%22light%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark_dimmed%22%2C%22color_mode%22%3A%22dark%22%7D%7D; _gh_sess=UMiKNOaMEpLxr9ue9DMPEojHtbFVD%2FICyB1WdezU0khXlxfBgnd3uKPJgTZxPowZeQ7Y70WmD3RCPm%2BonvZHGK%2FJtgYDFJtu%2BkFKB1AE%2BWYN4d2aK4cBr2sZVfxE%2B32aiRaOBjUP1IH0kNvPJ%2BZfRnrjFCNCbuYlDG3hB4YqkEsZaYdvkya69P8YBC1sTt0mbVXcDAHYgfpnTnjE%2BdbGmsy0zsI1UL6IvNf1d2Oj90DA7KMhk64mYj000iRLqtsHWh%2BXVkEvDO26yWdSyiFhWDb%2FKiMHdV4DpzgZ7dQdheNc4OWvkfevymx21uYo6H6EU0W56%2BxsOPVtRGkybWyjQVFig0pv8DJJWcL8rkBhyYg%3D--nWW%2BMEVs7WLYFcRh--MBYAZgXFa2MUMVyk%2FDVfGQ%3D%3D; preferred_color_mode=dark; tz=Asia%2FShanghai`
	m, _ := StringCookiesToMAP(cookies)
	for k, v := range m {
		fmt.Printf("%s:%s\n", k, v)
	}
}

func TestFlatten(t *testing.T) {
	jsonData := `{
						"first": "Dale",
						"last": "Murphy",
						"age": 44,
						"nets": ["ig", "fb", "value1", "value2"],
						"submap": {
							"subfirst": "Dale",
							"sublast": ["Murphy", "value3", "value4"],
							"subsub":{"subsub":["value5","value6","value7","value8","value9"]}
						}
				}`
	flattenedData, _ := FlattenJSON([]byte(jsonData))
	for _, ele := range flattenedData {
		fmt.Print(ele, "\n")
	}
}
