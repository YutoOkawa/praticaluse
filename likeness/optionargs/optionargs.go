package optionargs

type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛り
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// 麺の量、油揚げ、海老天の有無でインスタンス作成
// 一番愚直な方法
// func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
// 	return &Udon{
// 		men:      p,
// 		aburaage: aburaage,
// 		ebiten:   ebiten,
// 	}
// }

func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}

type Option struct {
	mem      Portion
	aburaage bool
	ebiten   uint
}

// 構造体を利用したオプション引数
// func NewUdon(opt Option) *Udon {
// 	// ゼロ値に対するデフォルト値処理は関数/メソッド内で行う
// 	// 朝食時間は海老天一本無料
// 	if opt.ebiten == 0 && time.Now().Hour() < 10 {
// 		opt.ebiten = 1
// 	}
// 	return &Udon{
// 		men:      opt.mem,
// 		aburaage: opt.aburaage,
// 		ebiten:   opt.ebiten,
// 	}
// }

// ビルダーを利用したオプション引数
// func NewUdon(p Portion) *fluentOpt {
// 	// デフォルトはコンストラクタ関数で指定
// 	return &fluentOpt{
// 		men:      p,
// 		aburaage: false,
// 		ebiten:   1,
// 	}
// }

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func useFluentInterface() {
	// oomoriKitsune := NewUdon(Large).Aburaage().Order()
}

type OptFunc func(r *Udon)

// Functional Optionsパターンを利用したオプション引数
func NewUdon(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) {
		r.men = p
	}
}

func OptAburaage() OptFunc {
	return func(r *Udon) {
		r.aburaage = true
	}
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) {
		r.ebiten = n
	}
}

func useFuncOption() {
	tokuseiUdon := NewUdon(OptAburaage(), OptEbiten(3))
}

func main() {
	// var tempuraUdon = NewUdon(Option{
	// 	mem:      Large,
	// 	aburaage: false,
	// 	ebiten:   1,
	// })
}
