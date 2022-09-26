package forms

type errors map[string][]string

// Add error Message for form fieled
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get Return index 0 message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
