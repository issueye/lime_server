package store

type Store interface {
	// 获取数据
	Get(key string) ([]byte, error)
	// 设置数据
	Set(key string, value []byte) error
	// 删除数据
	Del(key string) error
	// 关闭数据库
	Close() error
	// 遍历获取所有数据
	ForEach(prefix string, fn func(key string, value []byte) error) error
	// count 获取数据数量
	Count(prefix string) (int, error)
	// Clear 清空数据
	Clear(prefix string) error
}
