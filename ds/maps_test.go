package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Want an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertNoResult := func(t *testing.T, got string) {
		t.Helper()
		if got != "" {
			t.Fatal("got a result but didn't want one")
		}
	}

	dictionary := map[string]string{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		result, err := Search(dictionary, "test")
		assertNoError(t, err)
		assertStrings(t, result, "this is just a test")
	})

	t.Run("known word", func(t *testing.T) {
		result, err := Search(dictionary, "test1")
		assertError(t, err, ErrNotFound)
		assertNoResult(t, result)
	})
}
