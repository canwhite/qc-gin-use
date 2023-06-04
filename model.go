package main

import "gorm.io/gorm"

// model
type User struct {
	gorm.Model //会帮忙建id 、create and update time
	Name       string
	Age        uint
}

type Params struct {
	name  string
	email string
}

/**
CREATE TABLE menus (
  id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE sub_menus (
  id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  menu_id INT NOT NULL,
  FOREIGN KEY (menu_id) REFERENCES menus(id)
);

//如果正常写sql插入的话，是先创建一级菜单的数据
INSERT INTO menus (id, name) VALUES
(1, 'Menu A'),
(2, 'Menu B');
//最后，你需要插入二级菜单的数据，并指定对应的一级菜单的id

INSERT INTO sub_menus (id, name, menu_id) VALUES
(1, 'Sub Menu A1', 1),
(2, 'Sub Menu A2', 1),
(3, 'Sub Menu B1', 2),
(4, 'Sub Menu B2', 2);

//end ：用gorm写起来就相对简单点了，可以看sql中的实现



*/

// 定义一级菜单和二级菜单的机构体
type Menu struct {
	ID       int
	Name     string
	SubMenus []SubMenu `gorm:"foreignKey:MenuID;references:ID"`
}

type SubMenu struct {
	ID     int
	Name   string
	MenuID int
}
