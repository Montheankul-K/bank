package handler

import (
	"encoding/json"
	"github.com/Montheankul-K/bank/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{
		customerService: customerService,
	}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.customerService.GetCustomer(customerID)
	if err != nil {
		/*
			appErr, ok := err.(errs.AppError)
			if ok {
				w.WriteHeader(appErr.Code)
				fmt.Fprintln(w, appErr.Message)
			}
		*/
		handleError(w, err)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
