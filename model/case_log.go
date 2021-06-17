package model

type CaseLog struct {
	Level    string `json:"level"`
	Msg      string `json:"msg"`
	ReportId int    `json:"case_id"`
}

func AddCaseLog(log *CaseLog) bool {
	if err := db.Debug().Table("case_log").Create(&log).Error; err != nil {
		return false
	}
	return true
}
