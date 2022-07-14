package main

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template/*.mac
var Templates embed.FS

const (
	WRK_INPUT  = "input.txt"
	WRK_OUTPUT = "output.nwd"
	WRK_E3DLOG = "e3d.log"
	WRK_NAVLOG = "nav.log"
	WRK_BAT    = "makenwd.bat"
)

// MakePml makes a PML macro text from RvmBuilder
func MakePml(b *RvmBuilder) (string, error) {
	var buf bytes.Buffer
	if err := writeWithTemplate(&buf, "header.mac", map[string]any{
		"wd":  b.WorkingDir,
		"log": filepath.Join(b.WorkingDir, WRK_E3DLOG),
	}); err != nil {
		return "", err
	}
	for _, v := range b.ExportRvms {
		e := &ExportOption{
			RvmType:                "binary",
			Encoding:               "utfeight",
			ImpliedTube:            "separate",
			Holes:                  "on",
			Repr:                   "on",
			ReprCl:                 "off",
			ReprTube:               "on",
			ReprInsu:               "off",
			ReprObst:               "off",
			ReprLevel:              6,
			ReprLevelPipe:          6,
			ReprLevelNozz:          6,
			ReprLevelStru:          6,
			AutoColor:              "on",
			AutoColorDisplayExport: "on",
		}
		e.Merge(&b.ExportDefault)
		e.Merge(&v)
		if e.Filename == "" {
			return "", errors.New("RVM has no filename")
		}
		if err := writeWithTemplate(&buf, "exportrvm.mac", map[string]any{
			"file":   filepath.Join(b.WorkingDir, e.Filename),
			"text":   e.RvmType == "text",
			"note":   strings.ReplaceAll(e.Filenote, "'", "''"),
			"enc":    e.Encoding,
			"impl":   e.ImpliedTube,
			"holes":  e.Holes,
			"repr":   e.Repr,
			"cent":   e.ReprCl,
			"tube":   e.ReprTube,
			"insu":   e.ReprInsu,
			"obst":   e.ReprObst,
			"lv":     e.ReprLevel,
			"lvpipe": e.ReprLevelPipe,
			"lvnozz": e.ReprLevelNozz,
			"lvstru": e.ReprLevelStru,
			"ac":     e.AutoColor,
			"acdisp": e.AutoColorDisplayExport,
			"aclist": makeColorRule(e.AutoColorRules, e.GlobalTranslucency),
			"incl":   e.ExportInclude,
			"excl":   e.ExportExclude,
			"input":  filepath.Join(b.WorkingDir, WRK_INPUT),
		}); err != nil {
			return "", err
		}
	}
	if b.DumpAtt {
		if err := writeWithTemplate(&buf, "dumpatt.mac", map[string]any{
			"file":  filepath.Join(b.WorkingDir, b.DumpAttOption.Filename),
			"exatt": b.DumpAttOption.AdditionalAtts,
			"noval": b.DumpAttOption.HasNoData,
			"uda":   b.DumpAttOption.HasUda,
			"refs":  b.DumpAttOption.Selection,
		}); err != nil {
			return "", err
		}
	}
	if err := writeWithTemplate(&buf, "footer.mac", map[string]any{
		"bat":    filepath.Join(b.WorkingDir, WRK_BAT),
		"nav":    b.FiletoolsTaskRunner,
		"relin":  WRK_INPUT,
		"relout": WRK_OUTPUT,
		"rellog": WRK_NAVLOG,
		"ver":    b.NwdVersion,
		"out":    filepath.Join(b.WorkingDir, WRK_OUTPUT),
		"final":  b.Filename,
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func writeWithTemplate(w io.Writer, n string, d map[string]any) error {
	t, err := template.New(n).ParseFS(Templates, "template/"+n)
	if err != nil {
		return err
	}
	if err := t.Execute(w, d); err != nil {
		return err
	}
	return nil
}

func makeColorRule(acr []AutoColorRule, gt int) string {
	var buf bytes.Buffer
	cols := map[string]int{}
	for _, rule := range acr {
		hex := strings.ToLower(rule.Color)
		no, ok := cols[hex]
		if !ok {
			no = len(cols) + 1
			cols[hex] = no
			var r, g, b uint8
			// Scan hex color
			fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
			fmt.Fprintf(&buf, "colour %d rgb %d, %d, %d\n", no, r, g, b)
		}
		trans := gt
		if rule.Translucency != 0 {
			trans = rule.Translucency
		}
		buf.WriteString("export autocolour " + rule.Selection)
		fmt.Fprintf(&buf, " colour %d", no)
		if trans != 0 {
			fmt.Fprintf(&buf, " translucency %d", trans)
		}
		buf.WriteString("\n")
	}
	return strings.TrimSpace(buf.String())
}
