package hash

import (
	"crypto/sha1"
	"path/filepath"
	"github.com/sselph/scraper/testdata"
	"testing"
)

func TestSHA1(t *testing.T) {
	d, err := testdata.New()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	for _, f := range d.Files {
		if e := filepath.Ext(f.Path); !KnownExt(e) {
			t.Errorf("KnownExt(%q) => false; want true", e)
		}
		if got, err := Hash(f.Path, sha1.New()); err != nil {
			t.Errorf("Hash(%q, sha1.New()) => err = %v; want nil", f.Path, err)
		} else if got != f.SHA1 {
			t.Errorf("Hash(%q, sha1.New()) => %q; want %q", f.Path, got, f.SHA1)
		}
	}
}
