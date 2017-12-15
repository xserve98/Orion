package orion

import (
	"github.com/carousell/Orion/orion/handlers"
)

//RegisterEncoder allows for registering an HTTP request encoder to arbitrary urls
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterEncoder(svr Server, serviceName, method, httpMethod, path string, encoder handlers.Encoder) {
	if e, ok := svr.(handlers.Encodeable); ok {
		e.AddEncoder(serviceName, method, []string{httpMethod}, path, encoder)
	}
}

//RegisterEncoders allows for registering an HTTP request encoder to arbitrary urls
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterEncoders(svr Server, serviceName, method string, httpMethod []string, path string, encoder handlers.Encoder) {
	if e, ok := svr.(handlers.Encodeable); ok {
		e.AddEncoder(serviceName, method, httpMethod, path, encoder)
	}
}

//RegisterDecoder allows for registering an HTTP request decoder to a method
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterDecoder(svr Server, serviceName, method string, decoder handlers.Decoder) {
	if e, ok := svr.(handlers.Decodable); ok {
		e.AddDecoder(serviceName, method, decoder)
	}
}
