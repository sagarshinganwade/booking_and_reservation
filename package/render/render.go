package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sagarshinganwade/booking_and_reservation/package/config"
	"github.com/sagarshinganwade/booking_and_reservation/package/models"
)

var functions = template.FuncMap{}

// var tc = make(map[string]*template.Template)
var app *config.AppConfig

// NewTemplates sets config for template package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderPage renders a template
func RenderPage(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//check if requested page is available in cache
	// Create Template Cache
	//tc = app.TemplateCache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get tempaltes from Template Cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	// t.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creats a template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// fmt.Println("Inside createTemplateCache()")
	myCache := map[string]*template.Template{}

	//get all the files having name endinf with .page.html
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	// fmt.Println("pages: ", pages)
	if err != nil {
		return myCache, err
	}

	//range through pages

	for _, page := range pages {
		// fmt.Println("Page: ", page, "of Pages: ", pages)
		name := filepath.Base(page)
		// fmt.Println("name ", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		// fmt.Println("ts: ", *ts)
		if err != nil {
			return myCache, err
		}
		//log.Println("Values in TS: ", ts)
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			//log.Println("Values in TS+layout: ", ts)

		}
		myCache[name] = ts
	}
	// fmt.Println(myCache)
	return myCache, nil

}

// func RenderPage(w http.ResponseWriter, t string) {
// 	var temp *template.Template
// 	var err error

// 	//check for template in cache

// 	_, inMap := tc[t]
// 	if !inMap {
// 		//template is not in Map. Need to generate from storage
// 		log.Println("Fetching template from Storage")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 	} else {
// 		//template is in cache, fetch and write on browser
// 		log.Println("Using Cached Template.")

// 	}
// 	temp = tc[t]
// 	err = temp.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	//parse the template
// 	temp, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache

// 	tc[t] = temp

// 	return nil

// }
