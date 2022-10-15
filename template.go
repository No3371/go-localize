package main

import (
	"text/template"
)

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go-localize; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}

package {{ .Package }}

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var localizations = map[string]string{
{{- range $key, $element := .Localizations  }}
	"{{ $key }}":{{ call $.LineUp $key }} ` + "`{{ $element }}`," + `
{{- end }}
}

type Replacements map[string]interface{}

type Localizer struct {
	Locale         string
	FallbackLocale string
	Localizations  map[string]string
}

func New(locale string, fallbackLocale string) *Localizer {
	t := &Localizer{Locale: locale, FallbackLocale: fallbackLocale}
	t.Localizations = localizations
	return t
}

func (t Localizer) SetLocales(locale, fallback string) Localizer {
	t.Locale = locale
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) SetLocale(locale string) Localizer {
	t.Locale = locale
	return t
}

func (t Localizer) SetFallbackLocale(fallback string) Localizer {
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) GetWithLocale(locale, key string, replacements ...*Replacements) string {
	str, ok := t.Localizations[t.getLocalizationKey(locale, key)]
	if !ok {
		str, ok = t.Localizations[t.getLocalizationKey(t.FallbackLocale, key)]
		if !ok {
			return key
		}
	}

	// If the str doesn't have any substitutions, no need to
	// template.Execute.
	if strings.Index(str, "}}") == -1 {
		return str
	}

	return t.replace(str, replacements...)
}

func (t Localizer) Get(key string, replacements ...*Replacements) string {
	str := t.GetWithLocale(t.Locale, key, replacements...)
	return str
}

func GetWithLocale(locale, key string, replacements ...*Replacements) string {
	t := New(locale, locale)
	str := t.GetWithLocale(t.Locale, key, replacements...)
	return str
}

func (t Localizer) getLocalizationKey(locale string, key string) string {
	return fmt.Sprintf("%v.%v", locale, key)
}

func (t Localizer) replace(str string, replacements ...*Replacements) string {
	b := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return str
	}

	replacementsMerge := Replacements{}
	for _, replacement := range replacements {
		for k, v := range *replacement {
			replacementsMerge[k] = v
		}
	}

	err = template.Must(tmpl, err).Execute(b, replacementsMerge)
	if err != nil {
		return str
	}
	buff := b.String()
	return buff
}
`,
))
