package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var Wallet = struct {
	sync.RWMutex
	Balance float64
}{Balance: 235}

func CheckWalletHandler(w http.ResponseWriter, r *http.Request) {
	Wallet.RLock()
	fmt.Fprintf(w, "%v", Wallet.Balance)
	Wallet.RUnlock()
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	Wallet.Lock()
	if Wallet.Balance >= 20 {
		Wallet.Balance -= 20
		msg := fmt.Sprintf(
			"[%s] TRANSAKSI BERHASIL, SISA SALDO Rp.%v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			Wallet.Balance,
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%v", Wallet.Balance)))
	} else {
		msg := fmt.Sprintf(
			"[%s] SALDO TIDAK MENCUKUPI.\n",
			time.Now().Format("2006-01-02 15:04:05"),
		)
		SendMessageToWebSockets(msg)
		w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("%v", Wallet.Balance)))
	}
	Wallet.Unlock()
}
