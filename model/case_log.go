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

// 通过reportId查询日志
func QueryCaseLogByReportId(reportId int) (caseLogs []CaseLog) {
	if err := db.Debug().Table("case_log").Where("report_id = ?", reportId).Scan(&caseLogs).Error; err != nil {
		return nil
	}
	return
}
