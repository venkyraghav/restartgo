package maps

import "testing"

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)
	assertNoError(t, err)
	assertStrings(t, result, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Want an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertNoResult(t *testing.T, got string) {
	t.Helper()
	if got != "" {
		t.Fatal("got a result but didn't want one")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		result, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, result, "this is just a test")
	})

	t.Run("known word", func(t *testing.T) {
		result, err := dictionary.Search("test1")
		assertError(t, err, ErrNotFound)
		assertNoResult(t, result)
	})

	t.Run("Empty Key", func(t *testing.T) {
		result, err := dictionary.Search("")
		assertError(t, err, ErrKeyEmpty)
		assertNoResult(t, result)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
	t.Run("Empty key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("", "this is just a test")
		assertError(t, err, ErrKeyEmpty)
		result, err := dictionary.Search("")
		assertError(t, err, ErrKeyEmpty)
		assertNoResult(t, result)
	})
	t.Run("Empty value", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "")
		assertError(t, err, ErrValueEmpty)
	})
	t.Run("Word already exists", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "this is just a test")

		// Add the same word again
		err = dictionary.Add("test", "this is just a test")
		assertError(t, err, ErrKeyAlreadyExists)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("test")
		assertNoError(t, err)
		_, err = dictionary.Search("test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("Empty key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("")
		assertError(t, err, ErrKeyEmpty)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("Word does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "new definition")
		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "new definition")
	})

	t.Run("Empty key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("", "new definition")
		assertError(t, err, ErrKeyEmpty)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("Word does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new definition")
		assertError(t, err, ErrNotFound)
	})
}
