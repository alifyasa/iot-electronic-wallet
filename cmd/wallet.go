package main

import (
	"fmt"
	"net/http"
	"time"
)

var walletAmount = 235

func CheckWalletHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", walletAmount)
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	if walletAmount >= 20 {
		walletAmount = walletAmount - 20
		msg := fmt.Sprintf(
			"[%s] TRANSAKSI BERHASIL, SISA SALDO Rp.%v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			walletAmount,
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%v", walletAmount)))
	} else {
		msg := fmt.Sprintf(
			"[%s] SALDO TIDAK MENCUKUPI.\n",
			time.Now().Format("2006-01-02 15:04:05"),
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("%v", walletAmount)))
	}
}
