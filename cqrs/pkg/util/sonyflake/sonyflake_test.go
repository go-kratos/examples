package sonyflake

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	InitSonyflake()
}

func TestGenerateSonyflake(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := GenerateSonyflake()
		fmt.Println(id)
	}
}

func TestGenerateTime(t *testing.T) {
	// 秒
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)

	// 毫秒
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixMilli())
	fmt.Printf("时间戳（纳秒转换为毫秒）：%v;\n", time.Now().UnixNano()/1e6)

	// 微秒
	fmt.Printf("时间戳（微秒）：%v;\n", time.Now().UnixMicro())
	fmt.Printf("时间戳（纳秒转换为微秒）：%v;\n", time.Now().UnixNano()/1e3)

	// 纳秒
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
}
