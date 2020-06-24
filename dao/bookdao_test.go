package dao

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T)  {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试新增图书", testAddBooks)
	//t.Run("删除图书", testDeleteBookById)
}

//func testGetBooks(t *testing.T) {
//	books, _ := GetBooks()
//	for k,v := range books{
//		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
//	}
//}

//func testAddBooks(t *testing.T)  {
//	book := &model.Book{
//		Title:   "三国演义",
//		Author:  "罗贯中",
//		Price:   88.88,
//		Sales:   100,
//		Stock:   100,
//		ImgPath: "/static/img/default.jpg",
//	}
//	AddBook(book)
//}
//
//func testDeleteBookById(t *testing.T) {
//	DeleteBookById("30")
//}