package handlerfactory

type HuoBiFactory struct {
}

func (HBFac *HuoBiFactory) Create() BaseHandler {
	return &HuoBiHandler{}
}

type HuoBiHandler struct {}

func (HBF *HuoBiHandler) Get() {}

func (HBF *HuoBiHandler) Post() {}
