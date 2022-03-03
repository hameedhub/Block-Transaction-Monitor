package controller

import (
	"ERC1155/services"
	"ERC1155/util"
	"net/http"
	"strconv"
	"time"
)

type TransactionController interface {
	GetTransaction(http.ResponseWriter, *http.Request)
}

type transactionController struct {
	service services.BlocService
}

func NewBlockController( blockService services.BlocService) TransactionController {
	return transactionController{
		service : blockService,
	}
}

func (s transactionController) GetTransaction(w http.ResponseWriter,r *http.Request) {
	v := r.URL.Query()
	var rows int
	if v.Get("rows") != ""{
		row,err :=strconv.Atoi(v.Get("rows"))
		if err != nil {
			util.ErrorResponse(w, 400, 1, "Invalid rows supplied" )
			return
		}
		rows = row
	}
	dates := map[string]time.Time{}
	from_date :=  v.Get("from_date")
	if from_date != ""{
		format := "2006-01-02"
		from, err := time.Parse(format, from_date)
		if err != nil {
			util.ErrorResponse(w, 400, 1, "Invalid from_date supplied. e.g ("+format+")")
			return
		}
		to := time.Now().UTC()
		if v.Get("to_date") != ""{
		to, err = time.Parse(format, v.Get("to_date"))
			if err != nil {
				util.ErrorResponse(w, 400, 1, "Invalid to_date supplied. e.g ("+format+")")
				return
			}
		}

		dates = map[string]time.Time{ "from_date": from, "to_date":to }
	}else{
		dates = nil
	}

	filter := map[string]string{"from": v.Get("sender"), "to": v.Get("receiver") }

	d, _ := s.service.GetTransaction(filter, dates, rows)
	util.SuccessResponse(w,200, "success", d )
	return 
}