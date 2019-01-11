package rest

// Create a Method type
type Method int

const (
	GET Method = 1 + iota
	POST
	PUT
	DELETE
	PATCH
)

var method = [...]string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
}

func (m Method) String() string { return method[m-1] }
