package handlerfactory

type HuoBiFactory struct {
}

func (HBFac *HuoBiFactory) Create() BaseHandler {
	return &HuoBiHandler{}
}

type HuoBiHandler struct {
	HandlerMap map[string]func()
}

func (HBF *HuoBiHandler) Get() {}

func (HBF *HuoBiHandler) Post() {}
