package dao

import (
	"Goland_Mall/model"
	"Goland_Mall/utils"
)

//GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows,error := utils.DbHelper.Raw(sqlStr).Rows()
	if error != nil {
		return nil, error
	}
	var books []*model.Book //创建books切片
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}


////向数据库中新增一条
//func AddBook(b *model.Book) error {
//	//写sql语句
//	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
//	//执行
//	_, error := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
//	if error != nil {
//		return error
//	}
//	return nil
//}
//
//
////DeleteBook 根据图书的id从数据库中删除一本图书
//func DeleteBookById(bookId string) error {
//	//写sql语句
//	sqlStr := "delete from books where id = ?"
//	_, error := utils.Db.Exec(sqlStr, bookId)
//	if error != nil {
//		return error
//	}
//	return error
//}
//
//
////GetBookByID 根据图书的id从数据库中查询出一本图书
//func GetBookByID(bookId string) (*model.Book, error) {
//	//写sql语句
//	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
//	//执行
//	row := utils.Db.QueryRow(sqlStr, bookId)
//	//创建book对象
//	book := &model.Book{}
//	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
//	return book, nil
//}
//
//
////UpdateBook 根据图书的id更新图书信息
//func UpdateBook(b *model.Book) error {
//	//写sql语句
//	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
//	//执行
//	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
//	if err != nil {
//		return err
//	}
//	return nil
//}