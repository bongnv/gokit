package gokit

type Options struct {
	Path string
}

type handler struct {
	opts    *Options
	service *Service
}

// Generate ...
func Generate(opts Options) error {
	h := &handler{
		opts: &opts,
	}

	return h.do()
}

func (h *handler) do() error {
	if err := h.parseSource(); err != nil {
		return err
	}

	if err := h.generateEndpoints(); err != nil {
		return err
	}

	if err := h.generateServer(); err != nil {
		return err
	}

	return nil
}
