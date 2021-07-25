package contents

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Contents string
}
