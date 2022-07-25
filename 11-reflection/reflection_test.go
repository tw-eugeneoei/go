package reflection

import (
	"reflect"
	"testing"
)

type person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	t.Run("tests table", func(t *testing.T) {
		cases := []struct {
			name     string
			input    interface{}
			expected []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{
					"Chris",
				},
				[]string{"Chris"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{
					"Chris",
					"London",
				},
				[]string{"Chris", "London"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{
					"Chris",
					33,
				},
				[]string{"Chris"},
			},
			{
				"nested fields",
				person{
					"Chris",
					Profile{
						33,
						"London",
					},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointers to things",
				&person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			{
				"arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := []string{}
				Walk(c.input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, c.expected) {
					t.Errorf("got %v want %v", got, c.expected)
				}
			})
		}
	})

	t.Run("maps", func(t *testing.T) {
		inputMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		got := []string{}

		Walk(inputMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	contains := false
	for _, v := range haystack {
		if v == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %s but it didn't", haystack, needle)
	}
}
