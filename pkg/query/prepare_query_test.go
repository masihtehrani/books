package query

//type want struct {
//	args  []interface{}
//	query string
//}
//
//var cases = []struct {
//	entry Query
//	want  want
//}{
//	{
//		entry: Query{
//			Select: []string{"cards.id", "cards.pan_mask", "cards.full_name", "cards.bank"},
//			Body:   "FROM cards INNER JOIN user_cards ON cards.id = user_cards.card_id",
//			QueryFilters: structs.Query{
//				Filter: map[string]map[string][]string{
//					"bank": {"=": {"saman", "ansar"}},
//				},
//				Sort: map[string][]string{"asc": {"id"}},
//				Page: structs.Page{
//					Number: 1,
//					Size:   10,
//				},
//			},
//		},
//		want: want{
//			args:  []interface{}{"saman", "ansar"},
//			query: "SELECT cards.id, cards.pan_mask, cards.full_name, cards.bank FROM cards INNER JOIN user_cards ON cards.id = user_cards.card_id WHERE bank IN(?,?) ORDER BY id ASC  LIMIT 10  OFFSET 0",
//		},
//	},
//}
//
//func TestPrepareQuery(t *testing.T) {
//	for _, c := range cases {
//		q, err := New(nil)
//		if err != nil {
//			t.Error(err)
//		}
//
//		q.Select = c.entry.Select
//		q.Body = c.entry.Body
//		q.QueryFilters = c.entry.QueryFilters
//		q.wheres = c.entry.wheres
//		_ = q.prepareQuery()
//		got := q.Query
//
//		if c.want.query != got {
//			t.Errorf("expected %s but have %s", c.want.query, got)
//		}
//
//		for i := range c.want.args {
//			if q.args[i] != c.want.args[i] {
//				t.Errorf("got %s and want %s", q.args[i], c.want.args[i])
//				return
//			}
//		}
//	}
//}
