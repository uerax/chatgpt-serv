package db

func Ping() bool {
	d, err := db.DB()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	err = d.Ping()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	return true
}

// if table not exist `create table typeName (field)`
// else insert into table
func Create(tables any) error {
	return db.Create(tables).Error
}

// field为空全部插入, 不为空则插入指定的字段 
func InsertWith(tables any, fields ...string) error {
	if len(fields) == 0 {
		return Create(tables)
	}
	return db.Select(fields).Create(tables).Error
}

// field为空全部插入, 不为空则不插入指定的字段 
func InsertIgnore(tables any, fields ...string) error {
	if len(fields) == 0 {
		return Create(tables)
	}
	return db.Omit(fields...).Create(tables).Error
}

func InsertMap(table any, values any) error {
	return db.Model(table).Create(values).Error
}

// mod 0:主键升序 1:主键降序 其他:无指定
// table的值为查询的条件,无值无条件
func TakeOne(table any, mod int) error {
	switch mod {
	case 0:
		return db.First(table).Error
	case 1:
		return db.Last(table).Error
	default:
		return db.Take(table).Error
	}
}

// select * from table where cond
func Select(table any, cond any) error {
	if cond == nil {
		return db.Find(table).Error
	}
	return db.Where(cond).Find(table).Error
}

// query: name = ? and id = ?
// ?占位符
func SelectWithWhere(table any, query string, cond ...string) error {
	 return db.Where(query, cond).Find(table).Error
}

func SelectNotWhere(table any, query string, cond ...string) error {
	return db.Not(query, cond).Find(table).Error
}

func SelectOrWhere(table any, query string, cond ...string) error {
	return db.Or(query, cond).Find(table).Error
}

