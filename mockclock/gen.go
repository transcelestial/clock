package mockclock

//go:generate mockgen -source=../clock.go -destination=./clock.go -package=mockclock
//go:generate mockgen -source=../ticker.go -destination=./ticker.go -package=mockclock
//go:generate mockgen -source=../timer.go -destination=./timer.go -package=mockclock
