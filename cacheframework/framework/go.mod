module github.com/terrylin13/godemo/cacheframework/framework

go 1.17

require lru v0.0.0-00010101000000-000000000000 // indirect

require consistenthash v0.0.0-00010101000000-000000000000 // indirect

replace lru  v0.0.0=>./lru

replace consistenthash v0.0.0=> ./consistenthash