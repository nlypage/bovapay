package bovapay

type request struct {
	method   string
	endpoint string
	body     map[string]interface{}
}

func (r *request) Add(key string, value any) {
	if r.body == nil {
		r.body = make(map[string]interface{})
	}
	r.body[key] = value
}
