package main

import (
    "log"
    "strconv"

    "github.com/tealeg/xlsx"
)

func praseExcle(excelFileName string) ([]*UserInfo, error) {
    var userInfoList []*UserInfo
    var err error
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err == nil {
        if len(xlFile.Sheets) == 2 {
            for _, row := range xlFile.Sheets[1].Rows {
                if len(row.Cells) > 19 {
                    i, errPrase := strconv.ParseUint(row.Cells[0].Value, 10, 64)
                    if (errPrase == nil) && (i > 0) && len(row.Cells[1].Value) > 3 {
                        userInfoList = append(userInfoList,
                            &UserInfo{
                                No:                row.Cells[0].Value,
                                UserName:          row.Cells[1].Value,
                                Department:        row.Cells[2].Value,
                                Duty:              row.Cells[3].Value,
                                IDCard:            row.Cells[4].Value,
                                AnnualVacation:    row.Cells[5].Value,
                                CalculateSalary:   row.Cells[6].Value,
                                AttendanceDays:    row.Cells[7].Value,
                                AbsenceDeductions: row.Cells[8].Value,
                                TempIncDec:        row.Cells[9].Value,
                                ShouldPay:         row.Cells[10].Value,
                                Endowment:         row.Cells[11].Value,
                                Medical:           row.Cells[12].Value,
                                Unemployment:      row.Cells[13].Value,
                                Reserve:           row.Cells[14].Value,
                                PreTax:            row.Cells[15].Value,
                                DeductTax:         row.Cells[16].Value,
                                Fsalary:           row.Cells[17].Value,
                                Comment:           row.Cells[18].Value,
                                Email:             row.Cells[19].Value,
                            })
                    }
                }
            }
        } else {
            log.Println("Excel格式不对，请确认此Excel包含有两个Sheet！")
        }
    } else {
        log.Print(err.Error())

    }

    return userInfoList, err

}
