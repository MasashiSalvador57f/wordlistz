package model

import "time"

// Word is a entitie that represents record in table word in database.
type Word struct {
	ID        int64     `xorm:"bigint(20) unsigned not null pk 'id'"`
	Name      string    `xorm:"varchar(40) not null 'name'"`
	WordType  int64     `corm:"tinyint(3) not null 'word_type'"`
	Meaning   string    `xorm:"text not null 'meaning'"`
	CreatedAt time.Time `xorm:"datetime not null"`
	UpdatedAt time.Time `xorm:"datetime not null"`
}
