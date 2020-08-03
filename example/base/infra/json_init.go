// Code generated by 'freedom new-project base'
package infra

import (
	"github.com/8treenet/extjson"
	"github.com/8treenet/freedom/infra/requests"
)

func init() {
	// More references github.com/8treenet/extjson
	extjson.SetDefaultOption(extjson.ExtJSONEntityOption{
		NamedStyle:       extjson.NamedStyleLowerCamelCase,
		SliceNotNull:     true, // Empty slice does not return null, return []
		StructPtrNotNull: true, // Nil point does not return null, return {}
	})
	requests.Unmarshal = extjson.Unmarshal
	requests.Marshal = extjson.Marshal
}