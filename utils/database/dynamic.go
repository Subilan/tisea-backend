package database

import (
	"fmt"
	"tisea-backend/structs"
	"tisea-backend/utils/testing"
)

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

// 使用选取的列名和值获取对应的 *structs.DatabaseDynamic 数组。
//
// limit — 每次获取的数据量，offset 代表获取之前跳过的行数。
// 
// ☆ 一般而言，limit 即每次获取时需要的数据总量，offset 即先前已经获取的数据总量。
//
// ☆ 例如，若先前已经获取了 8 项，还要获取 7 项，则 limit=7, offset=8，在获取 7 项之前会跳过 8 项，实际获取的是第 9-15 项。
//
// desc — 是否按照 id 降序（descending）排列。
//
// ☆ 若为 false，则为按照升序（ascending）排列。降序时，id 最大的在前（最新）；升序时，id 最小的在前（最旧）
func GetDynamicsBy(valueName string, value interface{}, limit int, offset int, desc bool) ([]*structs.DatabaseDynamic, error) {
	orderBy := "DESC"

	if desc == false {
		orderBy = "ASC"
	}

	queryString := fmt.Sprintf("SELECT id, title, content, author, hidden, categories, tags, created_at, updated_at FROM `tisea_dynamics` WHERE %s=? ORDER BY id %s LIMIT %d OFFSET %d", valueName, orderBy, limit, offset)
	
	rows, err := Query(queryString, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*structs.DatabaseDynamic{}

	index := 0
	var current *structs.DatabaseDynamic
	for rows.Next() {
		current = new(structs.DatabaseDynamic)
		rows.Scan(&current.ID, &current.Title, &current.Content, &current.Author, &current.Hidden, &current.Categories, &current.Tags, &current.CreatedAt, &current.UpdatedAt)
		results = append(results, current)
		index++
	}

	return results, nil
}

// 通过作者获得动态的数组
func GetDynamicsByAuthor(author string, limit int, offset int, desc bool) ([]*structs.DatabaseDynamic, error) {
	return GetDynamicsBy("author", author, limit, offset, desc)
}

// 通过 ID 获得唯一对应的动态。若无对应动态，返回 nil
func GetDynamicsByID(id uint64) (*structs.DatabaseDynamic, error) {
	result, err := GetDynamicsBy("id", id, 1, 0, true)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result[0], err
}

// 通过 ID 删除唯一对应的动态
func DeleteDynamic(id uint64) error {
	execString := "DELETE FROM `tisea_dynamics` WHERE id=?"

	result, err := Exec(execString, id)

	if err != nil {
		return nil
	}

	if affected, affErr := result.RowsAffected(); affErr != nil {
		return nil
	} else {
		testing.MustEqual(result, affected)
	}

	return err
}