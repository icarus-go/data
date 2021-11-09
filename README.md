# Data

## 项目初衷 Remark

- 支持gorm.io的数据类型, 方便开发人员快速支持各种不同数据映射,省略在项目中的硬编码工作.

## 变更历史 History

- 主要支持MySQL的数据模型

## 现在 Now

> 目前版本: v0.0.1

- datetime 日期
    - 对应数据库 `MySQL`
    - 对应数据类型 `datetime`
- boolean
    - 对应数据库 `MySQL`
    - 对应数据类型 `tinyint(1)`
- date
    - 支持数据库 `MySQL`
    - 对应数据类型 `date`
- imodel
    - S
        - 字符串类型的主键表结构
        - 支持数据库 `Oracle,MySQL`
    - M
        - 整数型参数主键表结构
        - 主要支持雪花算法ID snowflakeID
        - 支持数据库 `MySQL`

## 未来 Todo

- 将支持Oracle的数据模型, 以及支持Oracle的驱动

## 使用 Using

```go
import (
    "pmo-test4.yz-intelligence.com/data"
)
```


