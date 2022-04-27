package service

type s struct {
	// a database dependency would go here but instead we're going to have a static map
	m map[int64]User
}

// New instantiates a new service.
func New( /* a database connection would be injected here */ ) *s {
	return &s{
		m: map[int64]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
			3: {ID: 3, Name: "Carol"},
		},
	}
}

func (s *s) GetUser(id int64) (result User, err error) {
	// instead of querying a database, we just query our static map
	if result, ok := s.m[id]; ok {
		return result, nil
	}

	return result, ErrNotFound
}

func (s *s) GetUsers(ids []int64) (result map[int64]User, err error) {
	// always a good idea to return non-nil maps to avoid nil pointer dereferences
	result = map[int64]User{}

	for _, id := range ids {
		if u, ok := s.m[id]; ok {
			result[id] = u
		}
	}

	return
}
