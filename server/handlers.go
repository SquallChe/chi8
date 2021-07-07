package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// GetPageCount handler
func GetPageCount(w http.ResponseWriter, r *http.Request) {

	connection := GetDBConnection()
	defer connection.Close()

	// 各値取得
	params := r.URL.Query()
	mode := params["mode"][0]

	jsonStr := ""
	strSearch := "where ValidFlag = 1 "
	var count int

	if mode == "1" {
		strSearch += "and RecommendFlag = 1"
	} else if mode == "2" {
		strSearch += fmt.Sprintf("and CategoryId = '%s' and KindId = '%s'", params["categoryId"][0], params["kindId"][0])
	} else if mode == "3" {
		strSearch += "and ("
		keyWords := strings.Replace(params["keyWord"][0], "　", " ", -1)

		for _, word := range strings.Split(keyWords, " ") {
			if word == "" {
				continue
			}
			strSearch += fmt.Sprintf("ItemName like '%%%s%%' or ", word)
		}

		strSearch = strings.TrimSuffix(strSearch, " or ") + ")"
	}

	sqlStr := fmt.Sprintf(`select count(*) from Item %s`, strSearch)

	if err := connection.QueryRow(sqlStr).Scan(&count); err != nil {
		log.Fatal(err.Error())
		return
	}

	if count%10 != 0 {
		count = count/10 + 1
	} else {
		count = count / 10
	}

	jsonStr = fmt.Sprintf(`{"count":"%d"}`, count)
	w.Write([]byte(jsonStr))
}

// GetItem handler
func GetItem(w http.ResponseWriter, r *http.Request) {

	connection := GetDBConnection()
	defer connection.Close()

	// 各値取得
	params := r.URL.Query()
	mode := params["mode"][0]
	page, _ := strconv.Atoi(params["page"][0])

	jsonStr := ""
	strSearch := "where ValidFlag = 1 "

	if mode == "1" {
		strSearch += "and RecommendFlag = 1"
	} else if mode == "2" {
		strSearch += fmt.Sprintf("and CategoryId = '%s' and KindId = '%s'", params["categoryId"][0], params["kindId"][0])
	} else if mode == "3" {
		strSearch += "and ("
		keyWords := strings.Replace(params["keyWord"][0], "　", " ", -1)

		for _, word := range strings.Split(keyWords, " ") {
			strSearch += fmt.Sprintf("ItemName like '%%%s%%' or ", word)
		}

		strSearch = strings.TrimSuffix(strSearch, " or ") + ")"
	}

	sqlStr := fmt.Sprintf(`
		select top 10 ItemId, ItemName, PriceJP, PriceCN, Description from (
		select row_number() over(order by ItemId) as id,
			ItemId, ItemName, PriceJP, PriceCN, Description
		from Item
		%s) s
		where s.id between 10 * (@page - 1) + 1 and 10 * @page`, strSearch)

	rows, err := connection.Query(sqlStr, sql.Named("page", page))

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(sqlStr)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var itemId,
			priceJP,
			priceCN int

		var itemName,
			description string

		if err := rows.Scan(&itemId, &itemName, &priceJP, &priceCN, &description); err != nil {
			return
		}
		jsonStr += fmt.Sprintf(`{"ItemId":"%d","ItemName":"%s","PriceJP":"%d","PriceCN":"%d","Description":"%s"},`, itemId, itemName, priceJP, priceCN, description)
	}

	jsonStr = fmt.Sprintf("[%s]", strings.TrimSuffix(jsonStr, ","))

	w.Write([]byte(jsonStr))
}

// GetDetailImages handler
func GetDetailById(w http.ResponseWriter, r *http.Request) {

	connection := GetDBConnection()
	defer connection.Close()

	// 各値取得
	params := r.URL.Query()
	itemId := params["itemId"][0]
	mode := params["mode"][0]
	keyWord := params["keyWord"][0]
	ip := r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
	jsonStr := ""

	folderPath := fmt.Sprintf(ITEM_IMAGE_FOLDER, itemId)
	files, _ := ioutil.ReadDir(folderPath)

	for _, file := range files {
		if !strings.Contains(file.Name(), "-") {
			continue
		}
		jsonStr += fmt.Sprintf(`{"src":"%s"},`, fmt.Sprintf(ITEM_IMAGE_PATH, itemId, file.Name()))
	}

	// get item info
	var itemName string
	var priceJP int
	if err := connection.QueryRow(`
		select ItemName, PriceJP from Item where ItemId = @itemId`,
		sql.Named("itemId", itemId)).Scan(&itemName, &priceJP); err != nil {
		log.Fatal(err.Error())
	}

	// save access log
	if _, err := connection.Exec(`
		insert into ItemAccessLog (ItemId,Mode,Keyword,IP)
		values (@itemId,@mode,@keyWord,@ip)`,
		sql.Named("itemId", itemId),
		sql.Named("mode", mode),
		sql.Named("keyWord", keyWord),
		sql.Named("ip", ip)); err != nil {
		log.Fatal(err.Error())
	}

	jsonStr = fmt.Sprintf(`{"imageItems":[%s],"itemName":"%s","priceJP":"%d"}`, strings.TrimSuffix(jsonStr, ","), itemName, priceJP)
	w.Write([]byte(jsonStr))
}

// GetKindByCategoryId handler
func GetKindByCategoryId(w http.ResponseWriter, r *http.Request) {

	connection := GetDBConnection()
	defer connection.Close()

	vars := mux.Vars(r)
	categoryId := vars["categoryId"]

	jsonStr := ""

	rows, err := connection.Query(`
        select KindId, KindName from KindMaster
		where CategoryId = @categoryId
		order by KindId
        --order by DisplaySeqno`,
		sql.Named("categoryId", categoryId))

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var kindId, kindName string
		if err := rows.Scan(&kindId, &kindName); err != nil {
			return
		}
		jsonStr += fmt.Sprintf(`{"kindId":"%s","kindName":"%s"},`, kindId, kindName)
	}

	jsonStr = fmt.Sprintf(`[%s]`, strings.TrimSuffix(jsonStr, ","))

	w.Write([]byte(jsonStr))
}

// AccessLog handler
func AccessLog(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("uid")
	fromName := r.FormValue("fromName")
	toName := r.FormValue("toName")

	content := fmt.Sprintf("「%s」から「%s」に遷移しました。", fromName, toName)
	WriteLog(fmt.Sprintf(LOG_FORMAT_NORMAL, userId, content))
}

// GetMenu handler
func GetMenu(w http.ResponseWriter, r *http.Request) {

	connection := GetDBConnection()
	defer connection.Close()

	//params := r.URL.Query()
	//roles := params["roles[]"]

	result := 0

	var errMsg,
		content,
		jsonStr string

	// 画面情報を取得
	strSql := `	select MenuId, MenuName, IconName, RouterName, ParamId
				from MenuMaster 
				where ValidFlag = 1
				order by SortId`

	if rows, err := connection.Query(strSql); err != nil {
		result = 9
		errMsg = err.Error()
		goto End
	} else {

		for rows.Next() {
			var menuId,
				menuName,
				iconName,
				routerName,
				paramId string
			if err := rows.Scan(&menuId, &menuName, &iconName, &routerName, &paramId); err != nil {
				result = 8
				errMsg = err.Error()
				goto End
			}

			content += fmt.Sprintf(`{"menuId":"%s","title":"%s","icon":"%s","link":"%s","id":"%s"},`, menuId, menuName, iconName, routerName, paramId)
		}
	}

	content = strings.TrimSuffix(content, ",")

End:

	if result != 0 {
		//WriteLog(fmt.Sprintf(LOG_FORMAT_ERROR, operatorId, "メニュー取得でエラーが発生しました。", result, errMsg))
	}

	jsonStr = fmt.Sprintf(`{"result":"%d","menu":[%s],"errMsg":"%s"}`, result, content, errMsg)
	w.Write([]byte(jsonStr))

}

// GetPageCount handler
func RegistMail(w http.ResponseWriter, r *http.Request) {

	// connection := GetDBConnection()
	// defer connection.Close()

	// 各値取得
	params := r.URL.Query()
	mail := params["mail"][0]

	jsonStr := ""
	result := SendMail(mail)

	jsonStr = fmt.Sprintf(`{"result":"%v"}`, result)
	w.Write([]byte(jsonStr))
}
