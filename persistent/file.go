package persistent

//文件接口
type File interface {
    getPrefix() string
    getSuffix() string
}
