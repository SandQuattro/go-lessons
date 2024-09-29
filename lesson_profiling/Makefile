.PHONY: cpuprof
cpuprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 -seconds 20 http://localhost:8080/debug/pprof/profile )

.PHONY: memprof
memprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 http://localhost:8080/debug/pprof/heap )

.PHONY: allocsprof
allocsprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 http://localhost:8080/debug/pprof/allocs )

.PHONY: goroutineprof
goroutineprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 http://localhost:8080/debug/pprof/goroutine )

.PHONY: blockprof
blockprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 http://localhost:8080/debug/pprof/block )

.PHONY: mutexprof
mutexprof:
	( PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8081 http://localhost:8080/debug/pprof/mutex )