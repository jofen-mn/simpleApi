package db

type User struct {
	ID         int    `grom:"id bigint(11) NOT NULL AUTO_INCREMENT"`
	Name       string `gorm:"name varchar(12) NOT NULL"`
	Number     string `gorm:"number varchar(18) NOT NULL"`
	Phone      string `gorm:"phone varchar(11) NOT NULL"`
	Gender     byte   `gorm:"gender smallint(2) NOT NULL"`
	Age        byte   `gorm:"age smallint(3) NOT NULL"`
	Stature    int16   `gorm:"stature smallint(3) NOT NULL"`
	Weight     int16    `gorm:"weight smallint(3) NOT NULL"`
	Address    string `gorm:"address varchar(30)"`
	Occupation string `gorm:"occupation varchar(10)"`
}

func (User) TableName() string {
	return "user"
}

type Relation struct {
	ID           int    `grom:"id bigint(11) NOT NULL AUTO_INCREMENT"`
	Origin       string `gorm:"origin varchar(18) NOT NULL"`
	Target       string `gorm:"target varchar(18) NOT NULL"`
	Relationship int8   `gorm:"relationship smallint(3) NOT NULL"`
}

func (Relation) TableName() string {
	return "relation"
}


