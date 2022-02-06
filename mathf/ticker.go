package mathf

type TickerImp interface {
	Tick(gameTime *GameTime)
	IsPaused() bool
	IsComplete() bool
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

// todo Remove(imp TickerImp)

func (t *Ticker) Tick(gameTime *GameTime) {
	var toRemove []TickerImp
	for _, imp := range t.tickers {
		if !imp.IsPaused() {
			imp.Tick(gameTime)
		}
		if imp.IsComplete() {
			toRemove = append(toRemove, imp)
		}
	}

	// remove any completed tickers
	if len(toRemove) > 0 {
		for i := len(t.tickers) - 1; i >= 0; i-- {
			for j := len(toRemove) - 1; j >= 0; j-- {
				if toRemove[j] == t.tickers[i] {
					end := len(t.tickers) - 1
					t.tickers[i] = t.tickers[end]
					t.tickers = t.tickers[:end]
					break
				}
			}
		}
	}
}
