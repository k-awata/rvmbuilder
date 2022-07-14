package main

import (
	"encoding/json"
	"io"
)

type AutoColorRule struct {
	Selection    string `json:"selection"`
	Color        string `json:"color"`
	Translucency int    `json:"translucency,omitempty"`
}

type DumpAttOption struct {
	Filename       string   `json:"filename"`
	Selection      string   `json:"selection"`
	AdditionalAtts []string `json:"additionalAtts"`
	HasNoData      bool     `json:"hasNoData"`
	HasUda         bool     `json:"hasUda"`
}

type ExportOption struct {
	Filename               string          `json:"filename,omitempty"`
	RvmType                string          `json:"rvmType,omitempty"`
	Filenote               string          `json:"filenote,omitempty"`
	Encoding               string          `json:"encoding,omitempty"`
	ImpliedTube            string          `json:"impliedTube,omitempty"`
	Holes                  string          `json:"holes,omitempty"`
	Repr                   string          `json:"repr,omitempty"`
	ReprCl                 string          `json:"reprCl,omitempty"`
	ReprTube               string          `json:"reprTube,omitempty"`
	ReprInsu               string          `json:"reprInsu,omitempty"`
	ReprObst               string          `json:"reprObst,omitempty"`
	ReprLevel              int             `json:"reprLevel,omitempty"`
	ReprLevelPipe          int             `json:"reprLevelPipe,omitempty"`
	ReprLevelNozz          int             `json:"reprLevelNozz,omitempty"`
	ReprLevelStru          int             `json:"reprLevelStru,omitempty"`
	AutoColor              string          `json:"autoColor,omitempty"`
	AutoColorDisplayExport string          `json:"autoColorDisplayExport,omitempty"`
	GlobalTranslucency     int             `json:"globalTranslucency,omitempty"`
	AutoColorRules         []AutoColorRule `json:"autoColorRules,omitempty"`
	ExportInclude          []string        `json:"exportInclude,omitempty"`
	ExportExclude          []string        `json:"exportExclude,omitempty"`
}

type RvmBuilder struct {
	Filename            string         `json:"filename"`
	WorkingDir          string         `json:"workingDir"`
	FiletoolsTaskRunner string         `json:"filetoolsTaskRunner"`
	NwdVersion          string         `json:"nwdVersion"`
	ExportDefault       ExportOption   `json:"exportDefault"`
	ExportRvms          []ExportOption `json:"exportRvms"`
	DumpAtt             bool           `json:"dumpAtt"`
	DumpAttOption       DumpAttOption  `json:"dumpAttOption"`
}

func LoadRvmBuilder(i io.Reader) (*RvmBuilder, error) {
	j, err := io.ReadAll(i)
	if err != nil {
		return nil, err
	}
	var b RvmBuilder
	if err := json.Unmarshal(j, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (e *ExportOption) Merge(o *ExportOption) {
	if o.Filename != "" {
		e.Filename = o.Filename
	}
	if o.RvmType != "" {
		e.RvmType = o.RvmType
	}
	if o.Filenote != "" {
		e.Filenote = o.Filenote
	}
	if o.Encoding != "" {
		e.Encoding = o.Encoding
	}
	if o.ImpliedTube != "" {
		e.ImpliedTube = o.ImpliedTube
	}
	if o.Holes != "" {
		e.Holes = o.Holes
	}
	if o.Repr != "" {
		e.Repr = o.Repr
	}
	if o.ReprCl != "" {
		e.ReprCl = o.ReprCl
	}
	if o.ReprTube != "" {
		e.ReprTube = o.ReprTube
	}
	if o.ReprInsu != "" {
		e.ReprInsu = o.ReprInsu
	}
	if o.ReprObst != "" {
		e.ReprObst = o.ReprObst
	}
	if o.ReprLevel != 0 {
		e.ReprLevel = o.ReprLevel
	}
	if o.ReprLevelPipe != 0 {
		e.ReprLevelPipe = o.ReprLevelPipe
	}
	if o.ReprLevelNozz != 0 {
		e.ReprLevelNozz = o.ReprLevelNozz
	}
	if o.ReprLevelStru != 0 {
		e.ReprLevelStru = o.ReprLevelStru
	}
	if o.AutoColor != "" {
		e.AutoColor = o.AutoColor
	}
	if o.AutoColorDisplayExport != "" {
		e.AutoColorDisplayExport = o.AutoColorDisplayExport
	}
	if o.GlobalTranslucency != 0 {
		e.GlobalTranslucency = o.GlobalTranslucency
	}
	if len(o.AutoColorRules) != 0 {
		e.AutoColorRules = o.AutoColorRules
	}
	if len(o.ExportInclude) != 0 {
		e.ExportInclude = o.ExportInclude
	}
	if len(o.ExportExclude) != 0 {
		e.ExportExclude = o.ExportExclude
	}
}
