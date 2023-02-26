package updater

const interfaceGo = `package json

// IsZeroer is an interface implemented by any type which wishes to
// convey whether its current value is the zero value. This is used by
// encoding/json and encoding/xml through its use of the "omitempty" option
// in struct tags.
type IsZeroer interface {
	IsZero() bool
}
`

var interfaceTestGo = []string{
	"package json",
	"",
	"import \"testing\"",
	"",
	`type neverzero int

func (nz neverzero) IsZero() bool {
	return false
}`,
	"",
	`type alwayszero int

func (az alwayszero) IsZero() bool {
	return true
}`,
	"",
	"type OptionalsIsZero struct {",
	"\tSr string `json:\"sr\"`",
	"\tSo string `json:\"so,omitempty\"`",
	"\tSw string `json:\"-\"`",
	"",
	"\tIr int `json:\"omitempty\"` // actually named omitempty, not an option",
	"\tIo int `json:\"io,omitempty\"`",
	"",
	"\tSlr []string `json:\"slr,random\"`",
	"\tSlo []string `json:\"slo,omitempty\"`",
	"",
	"\tMr map[string]any `json:\"mr\"`",
	"\tMo map[string]any `json:\",omitempty\"`",
	"",
	"\tFr float64 `json:\"fr\"`",
	"\tFo float64 `json:\"fo,omitempty\"`",
	"",
	"\tBr bool `json:\"br\"`",
	"\tBo bool `json:\"bo,omitempty\"`",
	"",
	"\tUr uint `json:\"ur\"`",
	"\tUo uint `json:\"uo,omitempty\"`",
	"",
	"\tStr struct{} `json:\"str\"`",
	"\tSto struct{} `json:\"sto,omitempty\"`",
	"",
	"\tNzr neverzero `json:\"nzr\"`",
	"\tNzo neverzero `json:\"nzo,omitempty\"`",
	"",
	"\tAzr alwayszero `json:\"azr\"`",
	"\tAzo alwayszero `json:\"azo,omitempty\"`",
	"",
	"\tPtrAzo *alwayszero `json:\"ptrAzo,omitempty\"`",
	"}",
	"",
	"var optionalsIsZeroExpected = `{",
	" \"sr\": \"\",",
	" \"omitempty\": 0,",
	" \"slr\": null,",
	" \"mr\": {},",
	" \"fr\": 0,",
	" \"br\": false,",
	" \"ur\": 0,",
	" \"str\": {},",
	" \"sto\": {},",
	" \"nzr\": 0,",
	" \"nzo\": 0,",
	" \"azr\": 0",
	"}`",
	"",
	`func TestOmitEmptyAndIsZero(t *testing.T) {
	var o OptionalsIsZero
	o.Sw = "something"
	o.Mr = map[string]any{}
	o.Mo = map[string]any{}

	got, err := MarshalIndent(&o, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	if got := string(got); got != optionalsIsZeroExpected {
		t.Errorf(" got: %s\nwant: %s\n", got, optionalsIsZeroExpected)
	}
}`,
	"",
}
