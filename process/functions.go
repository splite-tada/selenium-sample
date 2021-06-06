package process

import (
	"log"

	"github.com/sclevine/agouti"
	"github.com/yfrash/selenium-sample/config"
)

var driver *agouti.WebDriver
var page *agouti.Page

func initBrowser() {
	var err error

	// Driver起動
	driver = agouti.ChromeDriver()
	if err = driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}

	// ブラウザを起動
	page, err = driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
}

/*
 * ログイン処理
 */
func login(id string, password string) {
	// ログイン画面へ遷移
	if err := page.Navigate(config.GetUrl("login")); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	// email, passの要素を取得し、値を設定
	loginEmail := page.FindByID("email")
	loginPassword := page.FindByID("password")
	loginEmail.Fill(id)
	loginPassword.Fill(password)

	// サブミット
	page.FindByID("login-form").Submit()
}

/*
 * ログアウト処理
 */
func logout() {
	// マイページ画面へ遷移
	if err := page.Navigate(config.GetUrl("mypage")); err != nil {
		// TODO 遷移できなければログインしていない？
		return
	}

	page.FindByID("logout-form").Submit()
}
