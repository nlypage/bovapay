package bovapay

type authorization string

const (
	signatureAuthorization = "signature"
	authorizationToken     = "authorization_token"
)

type request struct {
	method            string
	endpoint          string
	body              map[string]interface{}
	authorizationType authorization
}

func (r *request) Add(key string, value any) {
	if r.body == nil {
		r.body = make(map[string]interface{})
	}
	r.body[key] = value
}
