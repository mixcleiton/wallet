package main

import "br.com.cleiton/wallet/internal/cmd"

func main() {
	startWallet := cmd.NewStartWallet()
	startWallet.StartWallet()
}
