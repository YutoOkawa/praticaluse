package tryiota

type CarType int

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

var t CarType

// t = SUV

const (
	a = iota // 0
	b        // 1
	c        // 2
	_        // 3だが使われない
	d        // 4
	e = iota // 5
)

const (
	f = iota // 再び0
	g        // 1
	h        // 2
)

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

var o CarOption

// o = SunRoof || HeatedSeat
