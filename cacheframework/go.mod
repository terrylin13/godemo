module github.com/terrylin13/godemo/cacheframework

go 1.17

replace framework => ./framework

replace lru => ./framework/lru

replace consistenthash => ./framework/consistenthash

require framework v0.0.0-00010101000000-000000000000

require (
	consistenthash v0.0.0-00010101000000-000000000000 // indirect
	lru v0.0.0-00010101000000-000000000000 // indirect
)
