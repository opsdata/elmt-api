package v1

import (
	"github.com/opsdata/common-base/pkg/scheme"
)

// GroupName is the group name used in API.
const GroupName = "api"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = scheme.GroupVersion{
	Group:   GroupName,
	Version: "v1",
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) scheme.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
