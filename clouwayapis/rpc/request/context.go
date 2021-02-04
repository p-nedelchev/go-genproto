package request

// ContextKey is an type to act as a Key of context values.
// Here is an example usage:
// var  contextAuthKey       = request.ContextKey("authorization")
type ContextKey string

// String implements Stringer interface.
func (c ContextKey) String() string {
	return "request " + string(c)
}
