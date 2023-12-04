package database

import "tisea-backend/structs"

func MakePostingDynamic(title string, content string, author string, categories string, tags string) *structs.PostingDynamic {
	dynamic := new(structs.PostingDynamic)

	dynamic.Title = title
	dynamic.Content = content
	dynamic.Author = author
	dynamic.Categories = categories
	dynamic.Tags = tags

	return dynamic
}

func InsertPostingDynamic(dynamic structs.PostingDynamic) error {
	_, err := Exec("INSERT INTO `tisea_dynamics` (title, content, author, categories, tags) VALUES (?, ?, ?, ?, ?)", dynamic.Title, dynamic.Content, dynamic.Author, dynamic.Categories, dynamic.Tags)

	return err
}

func GetDynamicsByAuthor(author string) ([]*structs.DatabaseDynamic, error) {
	queryString := "SELECT id, title, content, hidden, categories, tags, created_at, updated_at FROM `tisea_dynamics` WHERE author=?"
	
	rows, err := Query(queryString, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*structs.DatabaseDynamic{}

	index := 0
	var current *structs.DatabaseDynamic
	for rows.Next() {
		current = new(structs.DatabaseDynamic)
		rows.Scan(&current.Author, &current.Title, &current.Content, &current.Hidden, &current.Categories, &current.Tags, &current.CreatedAt, &current.UpdatedAt)
		results = append(results, current)
		index++
	}

	return results, nil
}