// +build mage

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	distPath   = "swag/swagger-ui-3.49.0/dist"
	docsPath   = "static/docs"
	exampleDef = "https://idratherbewriting.com/learnapidoc/docs/rest_api_specifications/openapi_openweathermap.yml"
	localDef   = "definition.yml"
	swagURL    = "https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.49.0.zip"
)

func Swagger() error {
	// setup
	if err := os.MkdirAll(filepath.Dir(docsPath), 0755); err != nil {
		return fmt.Errorf("cannot create docs path: %w", err)
	}

	// get swagger-ui source
	localSwagZip := "swag.zip"
	localSwagDir := "swag"
	if err := downloadFile(swagURL, localSwagZip); err != nil {
		return fmt.Errorf("cannot download swagger ui zip: %w", err)
	}
	if err := unzipSwaggerUI(localSwagZip, localSwagDir); err != nil {
		return fmt.Errorf("cannot unzip swagger ui file: %w", err)
	}

	// move dist to static docs
	if err := os.Rename(distPath, docsPath); err != nil {
		return fmt.Errorf("cannot move swagger ui dist to docs: %w", err)
	}

	// rewrite index doc contents
	if err := replaceDefinitionURL(docsPath); err != nil {
		return fmt.Errorf("cannot replace definition url: %w", err)
	}

	// download example definition file
	if err := downloadFile(exampleDef, path(docsPath, localDef)); err != nil {
		return fmt.Errorf("cannot download example definition: %w", err)
	}

	// cleanup
	if err := os.Remove(localSwagZip); err != nil {
		return fmt.Errorf("cannot remove swagger ui zip file: %w", err)
	}
	if err := os.RemoveAll(localSwagDir); err != nil {
		return fmt.Errorf("cannot remove swagger ui zip contents directory: %w", err)
	}

	return nil
}

func downloadFile(src, dest string) error {
	os.MkdirAll(filepath.Dir(dest), 0755)

	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %w", dest, err)
	}
	defer f.Close()

	resp, err := http.Get(src)
	if err != nil {
		return fmt.Errorf("cannot get url %s: %w", src, err)
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("cannot write response body to file: %w", err)
	}

	return nil
}

func unzipSwaggerUI(fname, dir string) error {
	r, err := zip.OpenReader(fname)
	if err != nil {
		return fmt.Errorf("cannot open zip reader: %w", err)
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("cannot open file %s: %w", f.Name, err)
		}
		defer rc.Close()

		fpath := filepath.Join(dir, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, f.Mode()); err != nil {
				return fmt.Errorf("cannot make directory %s: %w", fpath, err)
			}
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}
			if err := os.MkdirAll(fdir, f.Mode()); err != nil {
				return fmt.Errorf("cannot make directory %s: %w", fdir, err)
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return fmt.Errorf("cannot open file %s: %w", fpath, err)
			}
			defer f.Close()

			if _, err := io.Copy(f, rc); err != nil {
				return fmt.Errorf("cannot copy file %s: %w", f.Name(), err)
			}
		}
	}

	return nil
}

func replaceDefinitionURL(dir string) error {
	path := path(dir, "index.html")

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read file '%s': %w", path, err)
	}

	oldStr := "https://petstore.swagger.io/v2/swagger.json/definition.yml"
	contents := strings.Replace(string(b), oldStr, localDef, -1)

	if err := ioutil.WriteFile(path, []byte(contents), 0644); err != nil {
		return fmt.Errorf("cannot write file '%s': %w", path, err)
	}

	return nil
}

func path(parts ...string) string {
	return strings.Join(parts, string(os.PathSeparator))
}
