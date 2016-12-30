package mysql

import "testing"

func TestGetContext(t *testing.T) {
	// Precondition: a localhost MySQL connection at default port with:
	// * Test user test/test
	// * Test database golang_test
	// * A table context_test:
	//   * name (varchar(20))
	//   * quantity (int)
	// Both rows must be in the primary key.
	// Travis CI will have this set up already, see .travis.yml in the root.
	context, err := GetContext("test", "test", "", "golang_test")
	if err != nil {
		t.Fatalf("Unexpected error establishing context: %s", err)
	}

	columns, err := context.GetPrimaryKeyColumns("context_test")
	if err != nil {
		t.Fatalf("Unexpected error loading table: %s", err)
	}

	expectedColumns := []string{
		"name",
		"quantity",
	}

	for _, col := range expectedColumns {
		if _, ok := columns[col]; !ok {
			t.Fatalf("Column '%s' expected to be in primary key but not found", col)
		}
	}
}

func testGetGoTypeFromSqlType(t *testing.T) {
	cases := map[string]string{
		"varchar(20)": "string",
		"int(10)":     "int",
	}

	for input, expected := range cases {
		actual := getGoTypeFromSqlType(input)
		if actual != expected {
			t.Errorf("Expected '%s' to return '%s', but got '%s' instead", input, expected, actual)
		}
	}
}
