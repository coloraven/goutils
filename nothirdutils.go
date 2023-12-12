package goutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/maphash"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// CopyFile 复制文件
func CopyFile(sourceFile, destinationFile string) (err error) {
	input, err := os.ReadFile(sourceFile)
	if err != nil {
		return
	}

	err = os.WriteFile(destinationFile, input, 0644)
	return
}

// GetLocalIP 获取当前 IP
func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

// 根据下标删除 string slice 中的元素
func RemoveStringSliceItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// 判断两个 string slice 是否相同
func IsEqualStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, c := range s1 {
		if c != s2[i] {
			return false
		}
	}
	return true
}

// IsStrInSlice 判断字符串是否在给定的字符串切片中
func IsStrInSlice(i string, s []string) bool {
	if !sort.StringsAreSorted(s) {
		sort.Strings(s)
	}
	index := sort.SearchStrings(s, i)
	return (index < len(s) && s[index] == i)
}

// IsIntInSlice 判断整数是否在给定的整数切片中
func IsIntInSlice(i int, s []int) bool {
	if !sort.IntsAreSorted(s) {
		sort.Ints(s)
	}
	index := sort.SearchInts(s, i)
	return (index < len(s) && s[index] == i)
}

// ReverseFloat64Slice 反转 float64 切片，无返回值?
func ReverseFloat64Slice(numbers []float64) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

// ReversedFloat64Slice 反转 float64 切片，有返回值
func ReversedFloat64Slice(numbers []float64) []float64 {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// IntsMinMax 返回多个给定ints中的最小值和最大值
func IntsMinMax(ints ...int) (min int, max int) {
	if len(ints) == 0 {
		return
	}
	min = ints[0]
	max = ints[0]
	for _, value := range ints {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// RandInt 返回一个范围为 [start, end] 的int型随机数
func RandInt(start, end int) int {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	num := start + r.Intn(end-start+1)
	return num
}

// BytesToString 将 []byte 转换为字符串.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// CodePointToUTF8 converts Unicode Code Point to UTF-8.
// e.g.: CodePointToUTF8(`{"info":[["color","\u5496\u5561\u8272\u7c\u7eff\u8272"]]｝`, 16)
// => `{"info":[["color","咖啡色|绿色"]]｝`
func CodePointToUTF8(str string, base int) string {
	i := 0
	if strings.Index(str, `\u`) > 0 {
		i = 1
	}
	strSlice := strings.Split(str, `\u`)
	last := len(strSlice) - 1
	if len(strSlice[last]) > 4 {
		strSlice = append(strSlice, string(strSlice[last][4:]))
		strSlice[last] = string(strSlice[last][:4])
	}
	for ; i <= last; i++ {
		if x, err := strconv.ParseInt(strSlice[i], base, 32); err == nil {
			strSlice[i] = string(rune(x))
		}
	}
	return strings.Join(strSlice, "")
}

// FileExist 报告指定的文件或目录是否存在，并报告是否目录.
func FileExist(name string) (existed bool, isDir bool) {
	info, err := os.Stat(name)
	if err != nil {
		return !os.IsNotExist(err), false
	}
	return true, info.IsDir()
}

// 在指定路径中查找文件，返回文件全路径，如果文件不存在，返回错误.
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		fullpath = filepath.Join(path, filename)
		existed, _ := FileExist(fullpath)
		if existed {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

// 创建名字为 path 的目录,
// 可以指定必要的父目录，返回nil或者error
// The permission bits perm (before umask) are used for all
// directories that MkdirAll creates.
// 如果目录存在，则什么都不干，直接返回nil
// 如果权限代码为空, 则默认使用 0755.
func MkdirAll(path string, perm ...os.FileMode) error {
	var fm os.FileMode = 0755
	if len(perm) > 0 {
		fm = perm[0]
	}
	return os.MkdirAll(path, fm)
}

// WriteFile writes file, and automatically creates the directory if necessary.
// NOTE:
//
//	If perm is empty, automatically determine the file permissions based on extension.
func WriteFile(filename string, data []byte, perm ...os.FileMode) error {
	filename = filepath.FromSlash(filename)
	err := MkdirAll(filepath.Dir(filename))
	if err != nil {
		return err
	}
	if len(perm) > 0 {
		return os.WriteFile(filename, data, perm[0])
	}
	var ext string
	if idx := strings.LastIndex(filename, "."); idx != -1 {
		ext = filename[idx:]
	}
	switch ext {
	case ".sh", ".py", ".rb", ".bat", ".com", ".vbs", ".htm", ".run", ".App", ".exe", ".reg":
		return os.WriteFile(filename, data, 0755)
	default:
		return os.WriteFile(filename, data, 0644)
	}
}

// 格式化 JSON 数据并写入文件
func WriteJsonString(filename string, jsonBytes []byte) (bool, error) {
	// 创建一个用于接收格式化 JSON 的变量
	var out bytes.Buffer
	// 使用 json.Indent 格式化 JSON 数据
	err := json.Indent(&out, jsonBytes, "", "    ")
	if err != nil {
		return false, err
	}

	// 将格式化后的 JSON 数据写入文件
	err = os.WriteFile(filename, out.Bytes(), 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 传入日期字符串，返回格式化后的日期字符串
func GetDay(day string) (time.Time, error) {
	// 定义多个可能的日期时间布局
	layouts := []string{
		"2006-01-02",
		"20060102",
		"2006/01/02",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		time.RFC3339, // "2006-01-02T15:04:05Z07:00"
	}

	var parsedDate time.Time
	var err error

	// 尝试使用每种布局解析日期时间
	for _, layout := range layouts {
		parsedDate, err = time.Parse(layout, day)
		if err == nil {
			// 成功解析
			return parsedDate, nil
		}
	}

	// 如果所有布局都无法解析，则返回错误
	return time.Time{}, fmt.Errorf("unknown or invalid date format: %v", day)
}

func TraceTime() func() {
	pre := time.Now()
	return func() {
		elapsed := time.Since(pre)
		fmt.Println("Costs Time:\t", elapsed)
	}
}
