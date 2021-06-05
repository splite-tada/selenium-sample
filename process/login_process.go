package process

import (
	"selenium-sample/config"
	"time"
)

func LoginOnly() {
	initBrowser()

	// 処理後にDriver終了
	defer driver.Stop()

	login(config.GeneralLoginEmail, config.GeneralLoginPassword)

	time.Sleep(3 * time.Second)

	logout()

	time.Sleep(3 * time.Second)
}
