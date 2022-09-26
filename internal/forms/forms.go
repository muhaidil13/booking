package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Create new Form
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Required(field ...string) {
	for _, v := range field {
		value := f.Get(v)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(v, "Tolong Di isi Fieldnya")
		}
	}
}

// inisialisasi New Forms
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has Is Check Error Bagus digunakan pada checkbox radio button dll yang single karna
// parameternya bukan array atau veradic parameter
func (f *Form) Has(filed string) bool {
	x := f.Get(filed)
	if x == "" {
		f.Errors.Add(filed, "Tolong Isi Fieldnya dengan benar")
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This Field Must at Least %d character long", length))
		return false
	}
	return true
}

// Cek Email
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Addrees")
	}
}
