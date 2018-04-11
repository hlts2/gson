package gson

import (
	"testing"
)

func BenchmarkTestGetByKeys(b *testing.B) {
	json := `
		{"friends": [
			{
				"id": 0,
				"name": "a",
				"repositories": [
					{
						"id": 1,
						"title": "title1",
						"created_at": {
							"date": "2017-05-10 12:54:18",
               				"timezone": "UTC"
						}
					}
				]
			},
			{
				"id": 1,
				"name": "b",
				"repositories": [
					{
						"id": 1,
						"title": "title1",
						"created_at": {
							"date": "2017-05-10 12:54:18",
               				"timezone": "UTC"
						}
					}
				]
			}
		]}`

	g, _ := NewGsonFromByte([]byte(json))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := g.GetByKeys("friends", "0", "repositories", "0", "created_at", "timezone")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTestGetByPath(b *testing.B) {
	json := `
		{"friends": [
			{
				"id": 0,
				"name": "a",
				"repositories": [
					{
						"id": 1,
						"title": "title1",
						"created_at": {
							"date": "2017-05-10 12:54:18",
               				"timezone": "UTC"
						}
					}
				]
			},
			{
				"id": 1,
				"name": "b",
				"repositories": [
					{
						"id": 1,
						"title": "title1",
						"created_at": {
							"date": "2017-05-10 12:54:18",
               				"timezone": "UTC"
						}
					}
				]
			}
		]}`

	g, _ := NewGsonFromByte([]byte(json))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := g.GetByPath("friends.0.repositories.0.created_at.timezone")
		if err != nil {
			b.Fatal(err)
		}
	}
}
