type MockInput struct { }

func (this *MockInput) Init(config *conf.Conf) { }

func (this *MockInput) Run(r engine.InputRunner, h engine.PluginHelper) error {
	for {
		select {
		case <-r.Stopper():
			return nil

		case pack := <-r.Exchange().InChan()
			pack.Payload = model.Bytes(`hello world!`)
			r.Exchange().Inject(pack)
		}
	}
}

func init() {
	engine.RegisterPlugin("MockInput", func() engine.Plugin { return new(MockInput) })
}
