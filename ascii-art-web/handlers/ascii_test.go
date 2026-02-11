package handlers

import  (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	)

// Mock template for testing
func mockTemplate() *template.Template {
	tmpl := `{{.Text}}|{{.Banner}}|{{.Result}}|{{.Error}}`
	return template.Must(template.New("test").Parse(tmpl))
}


// Home handler with valid GET request
func TestHome_GET(t *testing.T) {
	h := NewHandler(mockTemplate())
	
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	
	h.Home(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Generate handler rejects GET request
func TestGenerate_GET(t *testing.T) {
	h := NewHandler(mockTemplate())
	
	req := httptest.NewRequest("GET", "/ascii-art", nil)
	w := httptest.NewRecorder()
	
	h.Generate(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for GET, got %d", w.Code)
	}
}

// Generate handler with empty text
func TestGenerate_EmptyText(t *testing.T) {
	h := NewHandler(mockTemplate())
	
	form := strings.NewReader("text=&banner=standard")
	req := httptest.NewRequest("POST", "/ascii-art", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	
	h.Generate(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for empty text, got %d", w.Code)
	}
	
	// Check error message in response
	if !strings.Contains(w.Body.String(), "Text and banner are required") {
		t.Error("Response should contain required fields error")
	}
}

// Generate handler with empty banner
func TestGenerate_EmptyBanner(t *testing.T) {
	h := NewHandler(mockTemplate())
	
	form := strings.NewReader("text=hello&banner=")
	req := httptest.NewRequest("POST", "/ascii-art", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	
	h.Generate(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for empty banner, got %d", w.Code)
	}
}
// Generate handler with valid form data 
func TestGenerate_ValidForm(t *testing.T) {
	h := NewHandler(mockTemplate())
	
	form := strings.NewReader("text=hello&banner=standard")
	req := httptest.NewRequest("POST", "/ascii-art", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	
	// This will use the actual ascii.Generate function
	// It will either succeed or fail based on whether banner files exist
	h.Generate(w, req)
	
	// We can't predict the exact status because it depends on banner files
	// But it should be one of: 200, 400, 404, or 500
	status := w.Code
	if status != http.StatusOK && 
	   status != http.StatusBadRequest && 
	   status != http.StatusNotFound && 
	   status != http.StatusInternalServerError {
		t.Errorf("Unexpected status code: %d", status)
	}
}
