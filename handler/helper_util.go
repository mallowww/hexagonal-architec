package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}

// appErr, ok := err.(errs.AppError)
// if ok {
// 	w.WriteHeader(appErr.Code)
// 	fmt.Fprintln(w, appErr.Message)
// 	return
// }
