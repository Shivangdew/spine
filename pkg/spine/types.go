package spine

type Ctx interface {
	Get(key string) (any, bool)
}
