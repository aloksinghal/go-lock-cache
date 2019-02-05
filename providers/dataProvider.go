package providers

type Provider interface {
	GetData(request map[string]interface{}) map[string] interface{}
}