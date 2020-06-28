package model

//Page 结构
type Page struct {
	Data        interface{} //每页查询出来的图书存放的切片
	CurrentPage int64       //当前页
	PageSize    int64       //每页显示的条数
	TotalPageNo int64       //总页数，通过计算得到
	TotalRecord int64       //总记录数，通过查询数据库得到
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.CurrentPage > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.CurrentPage < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.CurrentPage - 1
	}
	return 1
}

//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.CurrentPage + 1
	}
	return p.TotalPageNo
}

func OperatorData(data interface{},totalRecord,currentPage,pageSize int64) Page {
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	return Page{
		Data:       data,
		CurrentPage:   currentPage,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
}