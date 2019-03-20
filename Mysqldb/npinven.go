package mysqldbs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	NPINven "github.com/npinven/npinven"
)

type DocNoModel struct {
	DocNo         string `db:"DocNo"`
	DocDate       string `db:"DocDate"`
	MyDescription string `db:"MyDescription"`
	IsConfirm     int64  `db:"IsConfirm"`
	CreatorCode   string `db:"CreatorCode"`
	CreateDate    string `db:"CreateDateTime"`
	IsCancel      int64  `db:"ISCANCEL"`
}
type npinvenRepository struct{ db *sqlx.DB }

func NewNpinvenRepository(db *sqlx.DB) NPINven.Repository {
	return &npinvenRepository{db}
}

func (repo *npinvenRepository) GenDocNoInven(Type string, Search string, Branch string) (resp interface{}, err error) {
	now := time.Now()
	var lastdocno string

	var Docno string
	doc := DocNoModel{}
	sql := `select DocNo,DocDate,IsConfirm,CreatorCode,ISCANCEL from BCSTKInspect where CreatorCode = ? and DocNo like concat('%',?,'%')order by ROWORDER desc`
	err = repo.db.Get(&doc, sql, Search, Branch)
	if err != nil {
		return nil, err
	}
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
	fmt.Println(DocDate)
	fmt.Println(doc.DocDate[:10])
	if DocDate == doc.DocDate[:10] && doc.IsConfirm == 1 && doc.IsCancel == 0{
		return map[string]interface{}{
			"resp": map[string]interface{}{
				"isSuccess":   1,
				"processName": "Search Inspacet",
				"processDesc": "Successful",
				"data":        nil,
			},
			"docno": doc.DocNo}, nil

	} else {
		DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

		sql := `select DocNo,DocDate,IsConfirm,CreatorCode from BCSTKInspect where DocNo like concat('%',?,'%')order by ROWORDER desc`
		err = repo.db.Get(&doc, sql, Branch)
		if err != nil {
			return nil, err
		}
		docno, _ := strconv.Atoi(doc.DocNo[11:15])
		fmt.Println(doc.DocNo[6:10])
		if getDoctime() == doc.DocNo[6:10] {
			docno, _ = strconv.Atoi(doc.DocNo[11:15])
			fmt.Println(docno, "math")
		} else {
			docno = 0
		}
		last_number := (strconv.Itoa(docno + 1))
		fmt.Println(last_number)
		if len(string(last_number)) == 1 {
			lastdocno = "000" + string(last_number)
		}
		if len(string(last_number)) == 2 {
			lastdocno = "00" + string(last_number)
		}
		if len(string(last_number)) == 3 {
			lastdocno = "0" + string(last_number)
		}
		if len(string(last_number)) == 4 {
			lastdocno = string(last_number)
		}
		fmt.Println(lastdocno, "lastdoc")
		fmt.Println(getDoctime())
		Docno = Branch + "-IS" + getDoctime() + "-" + lastdocno
		fmt.Println(Docno)

		sqli := `insert into BCSTKInspect(DocNo,DocDate,MyDescription,InspectBy,IsConfirm,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?)`
		resp, err := repo.db.Exec(sqli, Docno, DocDate, "Mobile-app", 1, 1, Search, now.Format("2006-01-02 3:4:5"))
		if err != nil {
			fmt.Println("error = ", resp, err.Error())
			return nil, err
		}
		return map[string]interface{}{
			"resp": map[string]interface{}{
				"isSuccess":   1,
				"processName": "Search Inspacet",
				"processDesc": "Successful",
				"data":        nil,
			},
			"docno": Docno}, nil
	}

}

func getDoctime() string {
	var intyear int

	var vyear string

	var intmonth int
	var intmonth1 int
	var vmonth string
	var vmonth1 string
	var lenmonth int

	if time.Now().Year() >= 2560 {
		intyear = time.Now().Year()
	} else {
		intyear = time.Now().Year() + 543
	}

	vyear = strconv.Itoa(intyear)
	vyear1 := vyear[2:len(vyear)]

	fmt.Println("year = ", vyear1)

	intmonth = int(time.Now().Month())
	intmonth1 = int(intmonth)
	vmonth = strconv.Itoa(intmonth1)

	fmt.Println("month =", vmonth)

	lenmonth = len(vmonth)
	fmt.Println(vmonth1)
	if lenmonth == 1 {
		vmonth1 = "0" + vmonth
	} else {
		vmonth1 = vmonth
	}
	date := vyear1 + vmonth1
	fmt.Println(date)
	return date
}
