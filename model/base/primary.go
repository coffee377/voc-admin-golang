package base

/* 主键 */
type Primary struct {
	Id string `json:"id" orm:"column(ID);size(32);pk;description(主键ID)"` //主键ID
}

//func (p *Primary) SetId(id string) *Primary {
//	if id != "" {
//		p.Id = id
//	}
//	return p
//}
