package oasaSyncModel

type Sequences struct {
	SEQ_GEN   string `gorm:"primaryKey"`
	SEQ_COUNT int64
}
