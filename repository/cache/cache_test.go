package cache

import "testing"

func Test_WriteAndGet_Success(t *testing.T) {
	repo := New()
	//table-driven approach
	specs := []struct {
		key   string
		value string
		exp   string
	}{
		{
			key:   "from::to::word",
			value: "value",
			exp:   "value",
		},
		{
			key:   "::to::word",
			value: "",
			exp:   "",
		},
	}

	repo.Write(specs[0].key, specs[0].value)
	for _, spec := range specs {
		got := repo.Get(spec.key)
		if got != spec.exp {
			t.Fatalf("expected got to be %v but it is %v", spec.exp, got)
		}
	}

}
