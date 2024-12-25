package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MovieDetails struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"overview"`
}

func main() {
	http.HandleFunc("/get-movie-details", handleMovieDetails) // エンドポイントを設定

	fmt.Println("Server started at :8000")
	http.ListenAndServe(":8000", nil) // サーバーをポート8080で開始
}

// TMDB APIから映画情報を取得してフロントエンドに返すハンドラ
func handleMovieDetails(w http.ResponseWriter, r *http.Request) {
	// POSTリクエストを処理
	if r.Method == http.MethodPost {
		var requestData struct {
			AccountID int `json:"account_id"`
		}
		// フロントエンドから送られたデータを読み取る
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}

		// TMDB APIのURL（送られたIDを使用）
		url := fmt.Sprintf("https://api.themoviedb.org/3/account/%d", requestData.AccountID)

		// HTTPリクエストを作成
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		// ヘッダーの設定
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", "Bearer YOUR_API_KEY_HERE") // APIキーを設定

		// リクエストを実行
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// APIからのレスポンスを読み取る
		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		// レスポンスを構造体に変換
		var movie MovieDetails
		err = json.Unmarshal(body, &movie)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
			return
		}

		// フロントエンドにJSONレスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
