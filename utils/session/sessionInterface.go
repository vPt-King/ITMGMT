package utils

type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	SessionID() string
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}
