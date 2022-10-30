package mathf

type TickerImp interface {
	Resume()        // resumes ticking
	Pause()         // stops ticking but keeps state
	IsPaused() bool // if we are paused
	Tick()          // update our listeners
}

type Ticker struct {
	tickers []TickerImp
}

func NewTicker() *Ticker {
	return &Ticker{}
}

func (t *Ticker) Add(imp TickerImp) {
	t.tickers = append(t.tickers, imp)
}

func (t *Ticker) Tick() {
	for _, imp := range t.tickers {
		if !imp.IsPaused() {
			imp.Tick()
		}
	}
}
