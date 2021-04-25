package repository

//Cacher is implemented by objects that promote Cache repository
type Cacher interface {
	//Get returns value of cache with key argument
	Get(key string) string
	//Write inerts new key value in cache db
	Write(key, value string)
}
