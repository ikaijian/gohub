package sms

import "gohub/pkg/database"

func (sms *Sms) SaveSmsLog() {
	database.DB.Create(&sms)
}
