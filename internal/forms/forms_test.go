package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/testvalid", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("Tidak Valid")
	}
}

func TestRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/example", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Data Tidak Valid")
	}

	postData := url.Values{}
	// Menambah Data Dengan type key And value
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")
	r, err := http.NewRequest("POST", "/example", nil)
	if err != nil {
		t.Error("Data Kosong")
	}
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Data Tidak Valid")
	}
}

func TestMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	// Data Kosong
	postForm := r.PostForm
	post := New(postForm)
	post.MinLength("x", 10)
	// jika data Valid
	if post.Valid() {
		t.Error("data valid")
	}
	// mengambil nilai x dan jika tidak ada nilai error maka akan tampil error
	isError := post.Errors.Get("x")
	if isError == "" {
		t.Error("tidak Ada error")
	}

	// Test Valid
	postData := url.Values{}
	postData.Add("a", "123664")
	post = New(postData)
	post.MinLength("a", 100)
	// jika data valid maka akan error
	if post.Valid() {
		t.Error("Data A hanya sedikit sedangkan yang dibutuhkan itu banyak maka terjadi error")
	}

	// test not valid
	postData = url.Values{}
	postData.Add("b", "123664")
	post = New(postData)
	post.MinLength("b", 1)
	// Jika Tidak Valid
	if !post.Valid() {
		t.Error("Data Yang dimasukkan Sekarang Valid karena minimal panjang 1 sedangkan yang dimasukkan itu lebih")
	}
	isError = post.Errors.Get("b")
	if isError != "" {
		t.Error("jika data b ada error maka ini akan tampil")
	}

}

func TestIsEmail(t *testing.T) {
	postData := url.Values{}
	post := New(postData)
	post.IsEmail("Muha")
	if post.Valid() {
		t.Error("Bukan Email")
	}
	isError := post.Errors.Get("Muha")
	if isError == "" {
		t.Error("Ada Error")
	}
	// valid test
	postData = url.Values{}
	postData.Add("email", "muhaidil754@gmail.com")
	post = New(postData)
	post.IsEmail("email")
	if !post.Valid() {
		t.Error("Bukan Emial")
	}
}

func TestHas(t *testing.T) {
	postForm := url.Values{}
	formData := New(postForm)
	has := formData.Has("test")
	if has {
		t.Error("Data Ada")
	}
	values := url.Values{}
	values.Add("a", "a")
	formData = New(values)
	has = formData.Has("a")
	if !has {
		t.Error("Data Dicek dan Tidak Ada")
	}

}
