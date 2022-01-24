module github.com/terrylin13/godemo/cacheframework

go 1.17

replace framework => ./framework

replace lru => ./framework/lru

require framework v0.0.0-00010101000000-000000000000

require lru v0.0.0-00010101000000-000000000000 // indirect
