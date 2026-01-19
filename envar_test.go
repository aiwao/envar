package envar

import (
	"testing"
)

func setEnv(t *testing.T) {
	//error test
	t.Setenv("ENVAR_ERR", "erReRrerRoR")

	//int
	t.Setenv("ENVAR_8", "-14")
	t.Setenv("ENVAR_16", "-140")
	t.Setenv("ENVAR_32", "-140000")
	t.Setenv("ENVAR_64", "-9223372036854775140")
	
	//uint
	t.Setenv("ENVAR_U8", "14")
	t.Setenv("ENVAR_U16", "140")
	t.Setenv("ENVAR_U32", "140000")
	t.Setenv("ENVAR_U64", "9223372036854775140")

	//float
	t.Setenv("ENVAR_F32", "140.14")
	t.Setenv("ENVAR_F64", "1.797693134862315708145274237317043567981e+140")

	//complex
	t.Setenv("ENVAR_C64", "140i")
	t.Setenv("ENVAR_C128", "1400000000000000i")
	
	//byte
	t.Setenv("ENVAR_B", "14")

	//bool
	t.Setenv("ENVAR_BOOL", "true")

	//json
	t.Setenv("ENVAR_JSON", `{"env": "envar json"}`)

	//url
	t.Setenv("ENVAR_URL", "https://140.lol")
}

func TestInt(t *testing.T) {
	setEnv(t)
	i, err := GetInt("ENVAR_32")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetInt: %d", i)

	intValue := 12345
	if err := GetIntv("ENVAR_32", &intValue); err != nil {
		t.Fatal(err)
	}
	t.Logf("GetIntv: %d", intValue)

	//use default when error ocurred
	defaultValue := 12222
	GetIntv("ENVAR_ERR", &defaultValue)
	t.Logf("GetIntv (error): %d", defaultValue)
}

func TestBool(t *testing.T) {
	setEnv(t)
	boo, err := GetBool("ENVAR_BOOL")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetBool: %v", boo)
}

func TestUnmarshal(t *testing.T) {
	setEnv(t)
	var jsonValue struct{
		Env string `json:"env"`
	}
	if err := Unmarshal("ENVAR_JSON", &jsonValue); err != nil {
		t.Fatal(err)
	}
	t.Logf("Unmarshal: %s", jsonValue.Env)
}

func TestURL(t *testing.T) {
	setEnv(t)
	parsed, err := GetURL("ENVAR_URL")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("URL: %s", parsed.String())
}
