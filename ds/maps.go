package maps

type Dictionary map[string]string

type DictionaryErr string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrKeyEmpty         = DictionaryErr("key cannot be empty")
	ErrValueEmpty       = DictionaryErr("value cannot be empty")
	ErrKeyAlreadyExists = DictionaryErr("key already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func (d Dictionary) Search(key string) (string, error) {
	if isEmpty(key) {
		return "", ErrKeyEmpty
	}
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	if isEmpty(key) {
		return ErrKeyEmpty
	}
	if isEmpty(value) {
		return ErrValueEmpty
	}
	if _, ok := d[key]; ok {
		return ErrKeyAlreadyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	if isEmpty(key) {
		return ErrKeyEmpty
	}
	if _, ok := d[key]; !ok {
		return ErrNotFound
	}
	delete(d, key)
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if isEmpty(key) {
		return ErrKeyEmpty
	}
	if isEmpty(value) {
		return ErrValueEmpty
	}
	if _, ok := d[key]; !ok {
		return ErrNotFound
	}
	d[key] = value
	return nil
}
