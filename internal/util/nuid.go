package util

import (
	"crypto/rand"
	"math/big"
	"sync"
)

/*
* 薄雾算法
* 唯一ID
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | increase                                            | saltA    | saltB    |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子一，占 8 位，上限数值 255，使结果值不可预测；
* 3. 随机因子二，占 8 位，上限数值 255，使结果值不可预测；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

const saltBit = uint(8)                   // 随机因子 二进制位数
const saltShift = uint(8)                 // 随机因子 位移数
const increaseShift = saltBit + saltShift // 自增位移数

// NUid 结构
type NUid struct {
	sync.Mutex        // 互斥锁
	increase   uint64 // 自增数（实际47位）
	saltA      uint8  // 随机因子1
	saltB      uint8  // 随机因子2
}

// NewNUid 初始化
func NewNUid() *NUid {
	return &NUid{increase: 1}
}

// Generate 生成唯一ID
func (c *NUid) Generate() uint64 {
	c.Lock()
	c.increase++
	// 获取随机因子数值，真随机提高性能
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = uint8(randA.Int64())

	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = uint8(randB.Int64())
	// 占位，位运算
	nUid := (c.increase << increaseShift) | uint64(c.saltA)<<saltShift | uint64(c.saltB)
	c.Unlock()
	return nUid
}
