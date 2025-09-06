package encode

type OptionFlag uint16

const (
	HTMLEscapeOption    OptionFlag = 1 << iota //> html空格
	IndentOption                               //> 缩进
	UnorderedMapOption                         //> 非排序map
	DebugOption                                //> 调试
	ColorizeOption                             //> 彩色
	ContextOption                              //> with Context
	NormalizeUTF8Option                        //> utf8字符
	FieldQueryOption                           //> 字段查询
	FieldNoOmitEmpty                           //> 禁止忽略空值
	Int64ToStringOption                        //> int64(uint64)按照string处理
)
