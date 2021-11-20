package params

type (
	GetByID struct {
		// 2021-11-13 13:53:08 ps: 其实本意上是想使用雪花算法的
		// 无奈外包产品,有可能一台机器上多台MachineID(计算方式),所以放弃了,使用了 `auto_increment`
		ID uint64 `json:"id" form:"id" swaggertype:"string" example:"uint64 主键ID"` // 主键ID
	}

	GetByIDs struct {
		IDs []uint64 `json:"ids" form:"ids" swaggertype:"string" example:"[]uint64 主键IDs"` // 主键IDs
	}

	GetByUUID struct {
		ID string `json:"uuid" form:"uuid" example:"string UUID"` // ID UUID
	}

	GetByUUIDs struct {
		IDs []string `json:"uuids" form:"uuids" example:"[]string UUIDs"` // IDs UUIDs
	}
)
