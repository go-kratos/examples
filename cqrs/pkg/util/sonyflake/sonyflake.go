package sonyflake

import (
	"github.com/sony/sonyflake"
	"time"
)

var (
	sf *sonyflake.Sonyflake
)

func startTime() time.Time {
	return time.Now()
}

func getMachineID() (uint16, error) {
	return 0, nil
}

func checkMachineID(uint16) bool {
	return false
}

// InitSonyflake 初始化Sonyflake节点单体
func InitSonyflake() {
	settings := sonyflake.Settings{
		/*StartTime: startTime(),
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,*/
	}
	sf = sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("sonyflake not created")
	}
}

// GenerateSonyflake 生成 Sonyflake ID
func GenerateSonyflake() uint64 {
	if sf == nil {
		InitSonyflake()
	}
	if sf == nil {
		return 0
	}
	id, err := sf.NextID()
	if err != nil {
		return 0
	}
	return id
}
