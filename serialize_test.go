package bloom

import (
	"testing"
)

func TestJSON(t *testing.T) {
	s1 := "asöldkgjaösldkgaösldkasldgjkaösldkgjöasgkdjg"
	s2 := "elasödlnkgaölsdkfgaölsdkjfaölsdkgaölskgnaösl"
	s3 := "aölsdgkaösldkgaösldkgjaölsdkjgaölsdkgjaösldk"
	for n := 0; n < 100; n++ {
		for p := 1; p <= 128; p *= 2 {
			f1 := New(n, p)
			f1.Add(s1)
			f1.Add(s2)
			f1.Add(s3)
			bytes, encodingError := f1.EncodeJSON()
			if encodingError != nil {
				t.Errorf("f1.EncodeJSON() yields %v; want nil\n", encodingError)
			}
			f2, decodingError := DecodeJSON(bytes)
			if decodingError != nil {
				t.Errorf("f1.DecodeJSON() yields %v; want nil\n", decodingError)
			}
			if f2.lookups != f1.lookups {
				t.Errorf("f2.lookups = %v; want %v\n", f2.lookups, f1.lookups)
			}
			if f2.count != f1.count {
				t.Errorf("f2.count = %v; want %v\n", f2.count, f1.count)
			}
			if (f2.data == nil) != (f1.data == nil) {
				t.Errorf("f2.data = %v; want %v\n", f2.data, f1.data)
			}
			if len(f2.data) != len(f1.data) {
				t.Errorf("len(f2.data) = %v; want %v\n", len(f2.data), len(f1.data))
			}
			for i := range f2.data {
				if f2.data[i] != f1.data[i] {
					t.Errorf("f2.data[%v] = %v; want %v\n", i, f2.data[i], f1.data[i])
				}
			}
		}
	}
}
