package query

import "testing"

func TestSanitizeFilters(t *testing.T) {
	cases := []struct {
		name  string
		entry struct {
			selects []string
			filters map[string]map[string][]string
		}
		want map[string]map[string][]string
	}{
		{
			name: "success",
			entry: struct {
				selects []string
				filters map[string]map[string][]string
			}{
				selects: []string{"users.id", "card.pan_mask", "card.created_at"},
				filters: map[string]map[string][]string{
					"bank": {
						"eq": {
							"saman",
						},
					},
					"pan_mask": {
						"eq": {
							"6219*******37767653",
						},
					},
					"something": {
						"eq": {
							"somethingElse",
						},
					},
				},
			},
			want: map[string]map[string][]string{
				"bank": {
					"eq": {"saman"},
				},
				"pan_mask": {
					"eq": {
						"6219*******37767653",
					},
				},
			},
		},
	}

	q := Query{}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := q.sanitizeFilters(c.entry.selects, c.entry.filters)

			for key, value := range got {
				_, ok := c.want[key]
				if !ok {
					t.Errorf("not match for %s", c.want[key])
					return
				}

				for k, v := range value {
					if _, ok := c.want[key][k]; !ok {
						t.Errorf("not match for %s", c.want[key])
						return
					}
					for i := range c.want[key][k] {
						if v[i] != c.want[key][k][i] {
							t.Errorf("not match for %s", c.want[key])
							return
						}
					}
				}
			}
		})
	}
}
