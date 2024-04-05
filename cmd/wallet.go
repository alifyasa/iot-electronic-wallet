package main

import (
	"fmt"
	"net/http"
	"time"
)

var WalletBalance = 235

func CheckWalletHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", WalletBalance)
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	if WalletBalance >= 20 {
		WalletBalance = WalletBalance - 20
		msg := fmt.Sprintf(
			"[%s] TRANSAKSI BERHASIL, SISA SALDO Rp.%v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			WalletBalance,
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%v", WalletBalance)))
	} else {
		msg := fmt.Sprintf(
			"[%s] SALDO TIDAK MENCUKUPI.\n",
			time.Now().Format("2006-01-02 15:04:05"),
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("%v", WalletBalance)))
	}
}
