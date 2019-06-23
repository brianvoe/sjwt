package sjwt

import "testing"

func TestClaims(t *testing.T) {
	claims := New()
	claims.Add("temp", "temp val")
	claims.Add("bool", true)
	claims.Add("stringbool", "true")
	claims.Add("string", "hello world")
	claims.Add("intstring", 8675309)
	claims.Add("floatstring", 8675309.69)
	claims.Add("int", 8675309)
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
		t.Error("getting temp recieved incorrect value")
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
	floatstring := claims.GetStr("floatstring")
	if floatstring != "8675309.69" {
		t.Error("floatstring claim is incorrect, got: ", floatstring)
	}

	// Integer
	int := claims.GetInt("int")
	if int != 8675309 {
		t.Error("int claim is incorrect, got: ", int)
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
