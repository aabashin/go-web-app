package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no_file", true, "no error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no_file", true, "no error rendering non-existent jet template, when one is expected"},
	{"invalid_render_engine", "foo", "home", false, "no error rendering with non-existent template"},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRender.Renderer = e.renderer
		testRender.RootPath = "./testdata"

		err = testRender.Page(w, r, e.template, nil, nil)

		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}

	/*
			// цикл for выше позволяет избавиться от тестов в этой функции
		    r, err := http.NewRequest("GET", "/some-url", nil)
			if err != nil {
				t.Error(err)
			}

			w := httptest.NewRecorder()

			testRender.Renderer = "go"
			testRender.RootPath = "./testdata"

			err = testRender.Page(w, r, "home", nil, nil)

			if err != nil {
				t.Error("Error rendering page", err)
			}

			err = testRender.Page(w, r, "no-renderfile", nil, nil)

			if err == nil {
				t.Error("Error rendering go page", err)
			}

			testRender.Renderer = "jet"

			err = testRender.Page(w, r, "home", nil, nil)

			if err != nil {
				t.Error("Error rendering page", err)
			}

			err = testRender.Page(w, r, "no-renderfile", nil, nil)

			if err == nil {
				t.Error("Error rendering jet page", err)
			}

			testRender.Renderer = ""

			err = testRender.Page(w, r, "home", nil, nil)

			if err == nil {
				t.Error("No error returned while rendering with invalid renderer specified", err)
			}
	*/
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRender.Renderer = "go"
	testRender.RootPath = "./testdata"

	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRender.Renderer = "jet"

	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}
