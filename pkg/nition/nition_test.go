package nition

import (
	"testing"
)

func TestLoad(t *testing.T) {
	fset, err := LoadFingerprints()
	if err != nil {
		t.Errorf("LoadFingerprints() failed: %s", err)
	}
	if len(fset.Databases) == 0 {
		t.Errorf("LoadFingerprints() returned an empty set")
	}
}

func TestExamples(t *testing.T) {
	fset, err := LoadFingerprints()
	if err != nil {
		t.Errorf("LoadFingerprints() failed")
		return
	}
	if len(fset.Databases) == 0 {
		t.Errorf("LoadFingerprints() returned an empty set")
		return
	}
	for name, fdb := range fset.Databases {
		err := fdb.VerifyExamples()
		if err != nil {
			t.Errorf("VerifyExamples() failed for %s: %s", name, err)
		}
	}
}

func TestPJL(t *testing.T) {
	fset, err := LoadFingerprints()
	if err != nil {
		t.Errorf("LoadFingerprints() failed")
		return
	}
	if len(fset.Databases) == 0 {
		t.Errorf("LoadFingerprints() returned an empty set")
		return
	}

	m := fset.MatchFirst("hp_pjl_id.xml", "Xerox ColorQube 8570DT")
	if !m.Matched {
		t.Errorf("Failed to match 'Xerox ColorQube 8570DT': %#v", m)
		return
	}

	if m.Values["os.product"] != "8570DT" || m.Values["os.vendor"] != "Xerox" {
		t.Errorf("Failed to match 'Xerox ColorQube 8570DT' expected product or vendor")
	}
}

func TestPJLv2(t *testing.T) {
	fset, err := LoadFingerprints()
	if err != nil {
		t.Errorf("LoadFingerprints() failed")
		return
	}
	if len(fset.Databases) == 0 {
		t.Errorf("LoadFingerprints() returned an empty set")
		return
	}

	ms := fset.MatchAll("hp_pjl_id.xml", "Xerox ColorQube 8570DT")
	if len(ms) == 0 {
		t.Errorf("Failed to match 'Xerox ColorQube 8570DT'")
	}

	m := ms[0]

	if !m.Matched {
		t.Errorf("Failed to match 'Xerox ColorQube 8570DT'")
		return
	}

	if m.Values["os.product"] != "8570DT" || m.Values["os.vendor"] != "Xerox" {
		t.Errorf("Failed to match 'Xerox ColorQube 8570DT' expected product or vendor")
	}
}
