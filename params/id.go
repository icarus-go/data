package params

import "pmo-test4.yz-intelligence.com/kit/data/resource"

type (
	GetByID struct {
		// 2021-11-13 13:53:08 ps: 其实本意上是想使用雪花算法的
		// 无奈外包产品,有可能一台机器上多台MachineID(计算方式),所以放弃了,使用了 `auto_increment`
		ID uint64 `json:"id" form:"id" swaggertype:"string" example:"uint64 主键ID"` // 主键ID
	}

	//GetBySnowFlakeID 获取单条记录
	GetBySnowFlakeID struct {
		ID resource.SnowFlakeID `json:"id" swaggertype:"string" example:"string 雪花算法ID(实际为无符号整形)"` //  ID 主键
	}

	// GetBySnowFlakeIDs 获取ids
	GetBySnowFlakeIDs struct {
		IDs []resource.SnowFlakeID `json:"id" swaggertype:"string" example:"[]string 雪花算法ID(实际为无符号整形)"` //  ID 主键
	}

	GetByIDs struct {
		IDs []uint64 `json:"ids" form:"ids" swaggertype:"string" example:"[]uint64 主键IDs"` // 主键IDs
	}

	GetByUUID struct {
		ID string `json:"id" form:"id" example:"string UUID"` // ID UUID
	}

	GetByUUIDs struct {
		IDs []string `json:"ids" form:"ids" example:"[]string UUIDs"` // IDs UUIDs
	}
)
