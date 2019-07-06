package sjwt

import "testing"

func TestClaims(t *testing.T) {
	claims := New()
	claims.Add("temp", "temp val")
	claims.Add("bool", true)
	claims.Add("stringbool", "true")
	claims.Add("string", "hello world")
	claims.Add("intstring", 8675309)
	claims.Add("float32string", float32(86753.09))
	claims.Add("float64string", float64(86753.09))
	claims.Add("int", 8675309)
	claims.Add("uintint", uint(8675309))
	claims.Add("floatint", 86753.09)
	claims.Add("stringint", "8675309")
	claims.Add("float", 8675309.69)
	claims.Add("stringfloat", "8675309.69")

	// Check has function
	if !claims.Has("temp") {
		t.Error("temp doesnt exist when it should")
	}

	// Check normal get
	temp := claims.Get("temp").(string)
	if temp != "temp val" {
		t.Error("getting temp received incorrect value")
	}

	// Check deletion
	claims.Del("temp")
	if claims.Has("temp") {
		t.Error("temp exists when it should have been recently deleted")
	}

	// Boolean
	bool := claims.GetBool("bool")
	if bool != true {
		t.Error("bool claim is incorrect, got: ", bool)
	}
	stringbool := claims.GetBool("stringbool")
	if stringbool != true {
		t.Error("stringbool claim is incorrect, got: ", stringbool)
	}

	// String
	string := claims.GetStr("string")
	if string != "hello world" {
		t.Error("string claim is incorrect, got: ", string)
	}
	intstring := claims.GetStr("intstring")
	if intstring != "8675309" {
		t.Error("intstring claim is incorrect, got: ", intstring)
	}
	float32string := claims.GetStr("float32string")
	if float32string != "86753.09" {
		t.Error("float32string claim is incorrect, got: ", float32string)
	}
	float64string := claims.GetStr("float64string")
	if float64string != "86753.09" {
		t.Error("float64string claim is incorrect, got: ", float64string)
	}

	// Integer
	int := claims.GetInt("int")
	if int != 8675309 {
		t.Error("int claim is incorrect, got: ", int)
	}
	uintint := claims.GetInt("uintint")
	if uintint != 8675309 {
		t.Error("uintint claim is incorrect, got: ", uintint)
	}
	floatint := claims.GetInt("floatint")
	if floatint != 86753 {
		t.Error("floatint claim is incorrect, got: ", floatint)
	}
	stringint := claims.GetInt("stringint")
	if stringint != 8675309 {
		t.Error("stringint claim is incorrect, got: ", stringint)
	}

	// Float
	float := claims.GetFloat("float")
	if float != 8675309.69 {
		t.Error("float claim is incorrect, got: ", float)
	}
	stringfloat := claims.GetFloat("stringfloat")
	if stringfloat != 8675309.69 {
		t.Error("stringfloat claim is incorrect, got: ", stringfloat)
	}
}
