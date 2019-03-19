package mysqldbs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	NPINven "github.com/npinven/npinven"
)

type DocNoModel struct {
	DocNo       string `db:"DocNo"`
	DocDate     string `db:"DocDate"`
	IsConfirm   int64  `db:"IsConfirm"`
	CreatorCode string `db:"CreatorCode"`
}
type npinvenRepository struct{ db *sqlx.DB }

func NewNpinvenRepository(db *sqlx.DB) NPINven.Repository {
	return &npinvenRepository{db}
}

func (repo *npinvenRepository) GenDocNoInven(Type string, Search string, Branch string) (resp interface{}, err error) {
	now := time.Now()

	doc := DocNoModel{}
	sql := `select DocNo,DocDate,IsConfirm,CreatorCode from BCSTKInspect where CreatorCode = ? and DocNo like concat('%',?,'%')order by ROWORDER desc`
	err = repo.db.Get(&doc, sql, Search, Branch)
	if err != nil {
		return nil, err
	}
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
	fmt.Println(DocDate)
	fmt.Println(doc.DocDate[:10])
	if DocDate == doc.DocDate[:10] {
		return map[string]interface{}{
			"resp": map[string]interface{}{
				"isSuccess":   1,
				"processName": "Search Inspacet",
				"processDesc": "Successful",
				"data":        nil,
			},
			"docno": doc.DocNo}, nil

	} else {
		sql := `select DocNo,DocDate,IsConfirm,CreatorCode from BCSTKInspect where DocNo like concat('%',?,'%')order by ROWORDER desc`
		err = repo.db.Get(&doc, sql, Branch)
		if err != nil {
			return nil, err
		}
		docno, _ := strconv.Atoi(doc.DocNo[11:15])
		fmt.Println(len(string(docno+1)))

	}
	return map[string]interface{}{
		"resp": map[string]interface{}{
			"isSuccess":   1,
			"processName": "Search Inspacet",
			"processDesc": "Successful",
			"data":        nil,
		},
		"docno": doc.DocNo}, nil

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
