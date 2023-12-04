package goutils

import (
	"time"

	"github.com/coloraven/goutils"
	hc "github.com/feymanlee/holiday-cn"
)

type Holiday struct {
	Raw       string
	Time      time.Time
	IsHoliday bool
	IsOffDay  bool
	IsWeekend bool
	IsRestDay bool
}

func GetDayAttribut(day string) Holiday {
	var h Holiday
	h.Raw = day
	h.Time, _ = goutils.GetDay(day)
	h.IsOffDay = hc.IsOffDay(h.Time)
	h.IsHoliday = hc.IsHoliday(h.Time)
	//  是否是调休日
	h.IsRestDay = hc.IsRestDay(h.Time)
	h.IsWeekend = hc.IsWeekend(h.Time)
	return h
}

// type BaiduHolidays struct {
// 	DispEXT       interface{} `json:"DispExt"`
// 	QueryDispInfo interface{} `json:"QueryDispInfo"`
// 	ResultCode    int64       `json:"ResultCode"`
// 	ResultNum     int64       `json:"ResultNum"`
// 	QueryID       string      `json:"QueryID"`
// 	Result        []Result    `json:"Result"`
// }

// type Result struct {
// 	ClickNeed        string        `json:"ClickNeed"`
// 	DisplayData      DisplayData   `json:"DisplayData"`
// 	RecoverCacheTime string        `json:"RecoverCacheTime"`
// 	ResultURL        string        `json:"ResultURL"`
// 	Sort             string        `json:"Sort"`
// 	SrcID            string        `json:"SrcID"`
// 	SubResNum        string        `json:"SubResNum"`
// 	SubResult        []interface{} `json:"SubResult"`
// 	SuppInfo         string        `json:"SuppInfo"`
// 	Title            string        `json:"Title"`
// 	Weight           string        `json:"Weight"`
// 	ArPassthrough    ArPassthrough `json:"ar_passthrough"`
// }

// type ArPassthrough struct {
// 	OriginSrcid string `json:"origin_srcid"`
// 	TrueQuery   string `json:"true_query"`
// }

// type DisplayData struct {
// 	StdStg     string     `json:"StdStg"`
// 	StdStl     string     `json:"StdStl"`
// 	ResultData ResultData `json:"resultData"`
// 	Strategy   Strategy   `json:"strategy"`
// }

// type ResultData struct {
// 	EXTData EXTData `json:"extData"`
// 	TplData TplData `json:"tplData"`
// }

// type EXTData struct {
// 	OriginQuery string `json:"OriginQuery"`
// 	Clickneed   string `json:"clickneed"`
// 	Resourceid  string `json:"resourceid"`
// }

// type TplData struct {
// 	Data            Data            `json:"data"`
// 	DataSource      string          `json:"data_source"`
// 	DispDataURLEx   DispDataURLEx   `json:"disp_data_url_ex"`
// 	FeedbackContent FeedbackContent `json:"feedback_content"`
// 	GOtherinfo      GOtherinfo      `json:"g_otherinfo"`
// 	URLTransFeature URLTransFeature `json:"url_trans_feature"`
// }

// type Data struct {
// 	SiteID          int64     `json:"SiteId"`
// 	StdStg          int64     `json:"StdStg"`
// 	StdStl          int64     `json:"StdStl"`
// 	SelectTime      int64     `json:"_select_time"`
// 	UpdateTime      string    `json:"_update_time"`
// 	Version         int64     `json:"_version"`
// 	Almanac         []Almanac `json:"almanac"`
// 	AlmanacNumBaidu int64     `json:"almanac#num#baidu"`
// 	CambrianAppid   string    `json:"cambrian_appid"`
// 	Key             string    `json:"key"`
// 	LOC             string    `json:"loc"`
// 	Realurl         string    `json:"realurl"`
// 	Showlamp        string    `json:"showlamp"`
// 	URL             string    `json:"url"`
// 	XzhID           string    `json:"xzhId"`
// }

// type Almanac struct {
// 	Animal           Animal             `json:"animal"`
// 	Avoid            string             `json:"avoid"`
// 	CNDay            CNDay              `json:"cnDay"`
// 	Day              string             `json:"day"`
// 	FestivalInfoList []FestivalInfoList `json:"festivalInfoList,omitempty"`
// 	FestivalList     *string            `json:"festivalList,omitempty"`
// 	GzDate           string             `json:"gzDate"`
// 	GzMonth          GzMonth            `json:"gzMonth"`
// 	GzYear           GzYear             `json:"gzYear"`
// 	IsBigMonth       string             `json:"isBigMonth"`
// 	LDate            string             `json:"lDate"`
// 	LMonth           LMonth             `json:"lMonth"`
// 	LunarDate        string             `json:"lunarDate"`
// 	LunarMonth       string             `json:"lunarMonth"`
// 	LunarYear        string             `json:"lunarYear"`
// 	Month            string             `json:"month"`
// 	ODate            string             `json:"oDate"`
// 	Suit             string             `json:"suit"`
// 	Timestamp        string             `json:"timestamp"`
// 	Type             *string            `json:"type,omitempty"`
// 	Value            *string            `json:"value,omitempty"`
// 	Year             string             `json:"year"`
// 	YjJumpURL        string             `json:"yjJumpUrl"`
// 	YjFrom           YjFrom             `json:"yj_from"`
// 	Term             *string            `json:"term,omitempty"`
// 	Desc             *string            `json:"desc,omitempty"`
// 	Status           *string            `json:"status,omitempty"`
// }

// type FestivalInfoList struct {
// 	BaikeID   *string `json:"baikeId,omitempty"`
// 	BaikeName *string `json:"baikeName,omitempty"`
// 	BaikeURL  *string `json:"baikeUrl,omitempty"`
// 	Name      string  `json:"name"`
// }

// type DispDataURLEx struct {
// 	Aesplitid string `json:"aesplitid"`
// 	Sublink   string `json:"sublink"`
// 	SuppInfo  string `json:"suppInfo"`
// 	Suppinfo  string `json:"suppinfo"`
// 	Title     string `json:"title"`
// }

// type FeedbackContent struct {
// }

// type GOtherinfo struct {
// 	OriginQuery string `json:"OriginQuery"`
// 	IsShowBig   string `json:"is_show_big"`
// }

// type URLTransFeature struct {
// 	OriginQuery  string `json:"OriginQuery"`
// 	ClusterOrder string `json:"cluster_order"`
// 	IsShowBig    string `json:"is_show_big"`
// }

// type Strategy struct {
// 	CtplOrPHP   string `json:"ctplOrPhp"`
// 	HilightWord string `json:"hilightWord"`
// 	Precharge   string `json:"precharge"`
// }

// type Animal string

// const (
// 	兔 Animal = "兔"
// 	虎 Animal = "虎"
// )

// type CNDay string

// const (
// 	CNDay二 CNDay = "二"
// 	一      CNDay = "一"
// 	三      CNDay = "三"
// 	五      CNDay = "五"
// 	六      CNDay = "六"
// 	四      CNDay = "四"
// 	日      CNDay = "日"
// )

// type GzMonth string

// const (
// 	壬子 GzMonth = "壬子"
// 	甲寅 GzMonth = "甲寅"
// 	癸丑 GzMonth = "癸丑"
// 	辛亥 GzMonth = "辛亥"
// )

// type GzYear string

// const (
// 	壬寅 GzYear = "壬寅"
// 	癸卯 GzYear = "癸卯"
// )

// type LMonth string

// const (
// 	LMonth二 LMonth = "二"
// 	十一      LMonth = "十一"
// 	正       LMonth = "正"
// 	腊       LMonth = "腊"
// )

// type YjFrom string

// const (
// 	The51Wnl YjFrom = "51wnl"
// )

// func GetDayAttribut(year, month, day string) string {
// 	filePath := fmt.Sprintf("%s_%s.json", year, month) // 替换为您的 JSON 文件路径

// 	// 尝试打开文件
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		// 文件不存在或无法读取
// 		fmt.Println("File does not exist or cannot be read. Performing alternative operation.")
// 		jsonbyte := fetchHolidays(year, month)
// 		var respjson BaiduHolidays
// 		json.Unmarshal(jsonbyte, &respjson)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// fmt.Printf("%s\n", bodyText)
// 		// fmt.Printf("%s\n", )
// 		return respjson.Result[0].DisplayData.ResultData.TplData.Data.Almanac[0].Year
// 	} else {
// 		var respjson BaiduHolidays
// 		json.Unmarshal(data, &respjson)
// 	}
// }

// func fetchHolidays(year, month string) []byte {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", fmt.Sprintf("https://opendata.baidu.com/data/inner?tn=reserved_all_res_tn&type=json&resource_id=52109&query=%s年%s月&apiType=yearMonthData&cb", year, month), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Set("authority", "opendata.baidu.com")
// 	req.Header.Set("accept", "*/*")
// 	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-CN;q=0.8,en;q=0.7,zh-HK;q=0.6")
// 	// req.Header.Set("cache-control", "no-cache")
// 	// req.Header.Set("cookie", "BIDUPSID=13A4F2A8F2C711B7746826624FAF2242; PSTM=1686889222; BDUSS=W1kOGdWYUZOVmNWME1YQUZnVVQxRlpxekRFQkFqOEl3MnNRem95M3ZGZDlqYzlrSVFBQUFBJCQAAAAAAAAAAAEAAADY5joPc2lybGl1AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH0AqGR9AKhke; BDUSS_BFESS=W1kOGdWYUZOVmNWME1YQUZnVVQxRlpxekRFQkFqOEl3MnNRem95M3ZGZDlqYzlrSVFBQUFBJCQAAAAAAAAAAAEAAADY5joPc2lybGl1AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH0AqGR9AKhke; BAIDUID=13A4F2A8F2C711B7EA6C421EF42E67B1:SL=0:NR=10:FG=1; BAIDUID_BFESS=13A4F2A8F2C711B7EA6C421EF42E67B1:SL=0:NR=10:FG=1; ZFY=uEBQNjdxa:ASvO4vmHN8BbmDHJHdgbpTk:BXdX7tojrvQ:C; BDPASSGATE=IlPT2AEptyoA_yiU4SD93lIN8eDEUsCD34OtViBi3ECGh67BmhH74aJPE6r0NyerGTTMdI3JmMldjijsQmFuirMen46mfyJxbCHc-Mqfy1HmMsh3sLtoIab5VEE2sA8PbRhL-3MEF2ZEQlcFfObuhOw7ch_8bAlveeb55EvCpMrl0l7OYXn2xHyINnVkZmDAPNak-PrUd62ANpmfLvHwToPyaCI9QJ1Q7_mpg15XC3DkrkoXGAjQLxwHGIj748B9Bv_sHPaA81KH1i1z8Z2YSkUtdEiStMHdTiEVD_nYptJQTKHZQRW4VDrMHqcEjbPbLQdWKQ3zkt9UFi55vG9u0Z-8Qte9IWzGLSQyRNGGiBjZCXwVqTG3Cr3j9GVeKqvo0BR4KhpXTClZmFaAreCmpSr2HwPfbbNnT1Md_Si4u8gwfn63Gm4K6Hqze7BddneBMGGOG64H0qH8SIhKuuKyJaXaVKj8CCM26XPSgjamdj4o1fj6RMR7kzgCwnV-K7Gh6ru9VTDnVOO_eLZfwrXqvszDu9_U7lO3ujmILMy7Pfl518EcmqV4CmjJzGy_eWR2Zu_VtJX9FFDDnCdMwtyIjyZnlPY0AaJtC2j0JPITwcSIkWwV0yYlpMjDDBlDOpfMqtFKJ1TPGY8I3Tf9X7lRgrtA2V4vJYT6yNHK1oNPQD_r0DXaFMJU9EVKa7EtLeDZMVLipOsmYCGLk4AsP2qip7diFOElEjBzvXzJ3-T_4fzm94JHklkct0lUOYbmZur-8IQLoagqWyWn718-IbC5gjEB8YX6hRpzUam-Omvk8FKyIBEq-kW4; H_WISE_SIDS=110085_265881_275732_259642_281190_281867_281893_275095_280650_282196_282168_278414_282567_282630_282232_282848_253022_282402_283354_283223_251972_283364_283596_281704_283721_282888_256223_279610_283765_283867_283782_283896_284006_273240_284265_284264_283933_284575_265985_284690_284061_284217_283871_282607_284718_284794_284852_284879_284934_283795_282485_284810_285064_285179_285220_285277_285373_282427_285546_283903_285649_283659_285805_285818_285871_285872_285153_285876_281695_283014_277936_285992_284451_285765_256083_285938_278919_286192_281810_282174_286340_282466; PSINO=6; H_WISE_SIDS_BFESS=110085_265881_275732_259642_281190_281867_281893_275095_280650_282196_282168_278414_282567_282630_282232_282848_253022_282402_283354_283223_251972_283364_283596_281704_283721_282888_256223_279610_283765_283867_283782_283896_284006_273240_284265_284264_283933_284575_265985_284690_284061_284217_283871_282607_284718_284794_284852_284879_284934_283795_282485_284810_285064_285179_285220_285277_285373_282427_285546_283903_285649_283659_285805_285818_285871_285872_285153_285876_281695_283014_277936_285992_284451_285765_256083_285938_278919_286192_281810_282174_286340_282466; SE_LAUNCH=5%3A1701668979; BA_HECTOR=0ga40ka12l040l0ha00080231imqq3j1q; BDORZ=AE84CDB3A529C0F8A2B9DCDD1D18B695")
// 	// req.Header.Set("dnt", "1")
// 	// req.Header.Set("pragma", "no-cache")
// 	req.Header.Set("referer", "https://www.baidu.com/")
// 	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102"`)
// 	req.Header.Set("sec-ch-ua-mobile", "?0")
// 	req.Header.Set("sec-ch-ua-platform", "Windows")
// 	req.Header.Set("sec-fetch-dest", "script")
// 	req.Header.Set("sec-fetch-mode", "no-cors")
// 	req.Header.Set("sec-fetch-site", "same-site")
// 	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	bodyText, _ := io.ReadAll(resp.Body)
// 	goutils.WriteJsonString(fmt.Sprintf("%s_%s.json", year, month), bodyText)
// 	return bodyText
// }
