package gson

// FIXME(@hlts2): create benchmark test.
// var jsonString = `
// {"friends": [
// 	{
// 		"id": 0,
// 		"name": "a",
// 		"repositories": [
// 			{
// 				"id": 1,
// 				"title": "title1",
// 				"created_at": {
// 					"date": "2017-05-10 12:54:18",
//        				"timezone": "UTC"
// 				}
// 			}
// 		]
// 	},
// 	{
// 		"id": 1,
// 		"name": "b",
// 		"repositories": [
// 			{
// 				"id": 1,
// 				"title": "title1",
// 				"created_at": {
// 					"date": "2017-05-10 12:54:18",
//        				"timezone": "UTC"
// 				}
// 			}
// 		]
// 	}
// ]}`
//
// func BenchmarkTestGetByKeys(b *testing.B) {
// 	g, _ := NewGsonFromByte([]byte(jsonString))
//
// 	b.ResetTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		g.GetByKeys("friends", "0", "repositories", "0", "created_at", "timezone")
// 	}
// }
//
// func BenchmarkTestGetByPath(b *testing.B) {
// 	g, _ := NewGsonFromByte([]byte(jsonString))
//
// 	b.ResetTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		g.GetByPath("friends.0.repositories.0.created_at.timezone")
// 	}
// }
