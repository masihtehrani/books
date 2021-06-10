package query

import "testing"

func TestDotSeparate(t *testing.T) {
	cases := []struct {
		name  string
		entry string
		want  string
	}{
		{
			name:  "success",
			entry: "cards.id",
			want:  "id",
		}, {
			name:  "success",
			entry: "users.created_at",
			want:  "created_at",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := dotSeparate(c.entry)
			if c.want != got {
				t.Errorf("fail for %s : expected %s but get %s", c.entry, c.want, got)
			}
		})
	}
}
