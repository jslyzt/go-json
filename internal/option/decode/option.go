package decode

type OptionFlag uint8

const (
	FirstWinOption OptionFlag = 1 << iota
	ContextOption             //> with Context
	PathOption
	FieldNoOmitEmpty    //> 禁止忽略空值
	Int64ToStringOption //> int64(uint64)按照string处理
)
