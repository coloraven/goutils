# utils-go
 自己常用go语言函数


多个函数复制自：github.com/andeya/goutil


发现超强的武器库：
github.com/samber/lo
这个仓库如何食用

地址：https://gitee.com/hengy1/samber_example

    catalogue () 方法就是我都写了个 DEMO 以及简单说一下用法的思考
    recommend () 方法就是我挑选出来的 可能在开发中可能会常常使用到的地方
    samber 文件夹里面就是 从 a - z 像个目录一样查看就行
    在 mian.go 下面看到感兴趣的 然后详细就点进去看看备注以及我写的例子
```go
package main

import (
	"fmt"
	"strconv"
	"time"

	lop "github.com/samber/lo/parallel"
)

func main() {
	start := time.Now()
	lop.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		result := strconv.FormatInt(x, 10)
		time.Sleep(10 * time.Second)
		return result
	})
	fmt.Printf("耗时 %v\n", time.Since(start))
}
// []string{"1", "2", "3", "4"}
// 耗时 10.0080735s
```
## 节假日
```go
	t := goutils.GetDayAttribut("20230101")
	fmt.Println(t.IsHoliday, t.IsWeekend, t.IsOffDay, t.IsRestDay)
```