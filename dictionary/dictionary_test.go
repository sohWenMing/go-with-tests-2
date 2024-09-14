package dictionary

import (
	"fmt"
	"testing"
)

var dict = Dictionary{"test": "This is a dictionary test"}

/*
----------------------------------------
|                Tests                 |
----------------------------------------
*/
func TestSearch(t *testing.T) {

	t.Run("Testing Dictionary Search", func(t *testing.T) {

		lookup := "test"
		want := "This is a dictionary test"
		assertDefinition(t, dict, lookup, want)
	})
	t.Run("Testing failed lookup", func(t *testing.T) {
		lookup := "should fail"
		_, err := dict.Search(lookup)
		assertUnfoundError(t, err)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		key := "test-1"
		value := "this value should be added"
		err := dict.Add(key, value)
		assertNoError(t, err)
		assertDefinition(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this should no be added"
		err := dict.Add(key, value)
		assertDuplicateError(t, err)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("testing update, found key", func(t *testing.T) {
		key := "test"
		value := "the value should be updated"
		err := dict.Update(key, value)
		assertNoError(t, err)
		assertDefinition(t, dict, key, value)
	})
	t.Run("'testing update, key not found", func(t *testing.T) {
		key := "not found value"
		value := "the value should not be updated"
		err := dict.Update(key, value)
		assertUnfoundError(t, err)

	})
}

func TestErrorGenerator(t *testing.T) {

	const key = "test value"

	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "Test errNotFound",
			got:  generateError(errNotFound, key).Error(),
			want: fmt.Errorf("could not find the word you were looking for : %s", key).Error(),
		}, {
			name: "Test errWordExists",
			got:  generateError(errWordExists, key).Error(),
			want: fmt.Errorf("cannot add the word : %s because it already exists in the dictionary", key).Error(),
		},
		{
			name: "Test errNotFoundOnUpdate",
			got:  generateError(errNotFoundOnUpdate, key).Error(),
			want: fmt.Errorf("could not find the word you were looking for to update: %s", key).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkGeneratedError(t, test.got, test.want)
		})

	}
}

/*
----------------------------------------
|                Helpers                |
----------------------------------------
*/
func assertDefinition(t testing.TB, dict Dictionary, lookup, want string) {
	t.Helper()
	got, err := dict.Search(lookup)
	assertNoError(t, err)
	if got != want {
		t.Error(generateErrorString(got, want))
	}

}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("unexpected error occured.")
	}
}

func assertUnfoundError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error, but didn't get one")
	}
}

func assertDuplicateError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("wanted an error, but didn't get one")
	}
}

func checkGeneratedError(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Error(generateErrorString(got, want))
	}
}

func generateErrorString(got, want string) string {
	return fmt.Sprintf("got: %s\n want: %s \n", got, want)
}
