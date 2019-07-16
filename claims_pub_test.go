package sjwt

import "testing"

func TestClaims(t *testing.T) {
	claims := New()
	claims.Set("temp", "temp val")
	claims.Set("bool", true)
	claims.Set("stringbool", "true")
	claims.Set("string", "hello world")
	claims.Set("intstring", 8675309)
	claims.Set("float32string", float32(86753.09))
	claims.Set("float64string", float64(86753.09))
	claims.Set("int", 8675309)
	claims.Set("uintint", uint(8675309))
	claims.Set("floatint", 86753.09)
	claims.Set("stringint", "8675309")
	claims.Set("float", 8675309.69)
	claims.Set("stringfloat", "8675309.69")

	// Check has function
	if !claims.Has("temp") {
		t.Error("temp doesnt exist when it should")
	}

	// Check normal get
	temp, _ := claims.Get("temp")
	if temp.(string) != "temp val" {
		t.Error("getting temp received incorrect value")
	}

	// Check deletion
	claims.Del("temp")
	if claims.Has("temp") {
		t.Error("temp exists when it should have been recently deleted")
	}

	// Boolean
	bool, _ := claims.GetBool("bool")
	if bool != true {
		t.Error("bool claim is incorrect, got: ", bool)
	}
	stringbool, _ := claims.GetBool("stringbool")
	if stringbool != true {
		t.Error("stringbool claim is incorrect, got: ", stringbool)
	}

	// String
	string, _ := claims.GetStr("string")
	if string != "hello world" {
		t.Error("string claim is incorrect, got: ", string)
	}
	intstring, _ := claims.GetStr("intstring")
	if intstring != "8675309" {
		t.Error("intstring claim is incorrect, got: ", intstring)
	}
	float32string, _ := claims.GetStr("float32string")
	if float32string != "86753.09" {
		t.Error("float32string claim is incorrect, got: ", float32string)
	}
	float64string, _ := claims.GetStr("float64string")
	if float64string != "86753.09" {
		t.Error("float64string claim is incorrect, got: ", float64string)
	}

	// Integer
	int, _ := claims.GetInt("int")
	if int != 8675309 {
		t.Error("int claim is incorrect, got: ", int)
	}
	uintint, _ := claims.GetInt("uintint")
	if uintint != 8675309 {
		t.Error("uintint claim is incorrect, got: ", uintint)
	}
	floatint, _ := claims.GetInt("floatint")
	if floatint != 86753 {
		t.Error("floatint claim is incorrect, got: ", floatint)
	}
	stringint, _ := claims.GetInt("stringint")
	if stringint != 8675309 {
		t.Error("stringint claim is incorrect, got: ", stringint)
	}

	// Float
	float, _ := claims.GetFloat("float")
	if float != 8675309.69 {
		t.Error("float claim is incorrect, got: ", float)
	}
	stringfloat, _ := claims.GetFloat("stringfloat")
	if stringfloat != 8675309.69 {
		t.Error("stringfloat claim is incorrect, got: ", stringfloat)
	}
}
