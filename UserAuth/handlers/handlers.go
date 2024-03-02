package handlers

import (
	db "UserAuth/database"
	jt "UserAuth/json_token"
	jwtT "UserAuth/jwt_token"
	sl "UserAuth/serverLog"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Guid struct {
	Guid string `json:"guid"`
}

type MsgJson struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

var IndexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"███╗░░░███╗███████╗██████╗░░█████╗░██████╗░░██████╗  ░█████╗░░█████╗░███╗░░░███╗██████╗░░█████╗░███╗░░██╗██╗░░░██╗  ░██████╗███████╗██████╗░██╗░░░██╗███████╗██████╗░\n"+
			"████╗░████║██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝  ██╔══██╗██╔══██╗████╗░████║██╔══██╗██╔══██╗████╗░██║╚██╗░██╔╝  ██╔════╝██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗\n"+
			"██╔████╔██║█████╗░░██║░░██║██║░░██║██║░░██║╚█████╗░  ██║░░╚═╝██║░░██║██╔████╔██║██████╔╝███████║██╔██╗██║░╚████╔╝░  ╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝█████╗░░██████╔╝\n"+
			"██║╚██╔╝██║██╔══╝░░██║░░██║██║░░██║██║░░██║░╚═══██╗  ██║░░██╗██║░░██║██║╚██╔╝██║██╔═══╝░██╔══██║██║╚████║░░╚██╔╝░░  ░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██╔══╝░░██╔══██╗\n"+
			"██║░╚═╝░██║███████╗██████╔╝╚█████╔╝██████╔╝██████╔╝  ╚█████╔╝╚█████╔╝██║░╚═╝░██║██║░░░░░██║░░██║██║░╚███║░░░██║░░░  ██████╔╝███████╗██║░░██║░░╚██╔╝░░███████╗██║░░██║\n"+
			"╚═╝░░░░░╚═╝╚══════╝╚═════╝░░╚════╝░╚═════╝░╚═════╝░  ░╚════╝░░╚════╝░╚═╝░░░░░╚═╝╚═╝░░░░░╚═╝░░╚═╝╚═╝░░╚══╝░░░╚═╝░░░  ╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝\n\n"+
			"█▀▀ █▀█ █▀▀ ▄▀█ ▀█▀ █▀▀ █▀▄   █▄▄ █▄█   ▄▀█ █▄░█ █░█ ▄▀█ █▀█   █▀▀ ▄▀█ █▀█ █▀▀ █▀▀ █░█   █ █▄░█   ▀█ █▀█ ▀█ █░█\n"+
			"█▄▄ █▀▄ ██▄ █▀█ ░█░ ██▄ █▄▀   █▄█ ░█░   █▀█ █░▀█ ▀▄▀ █▀█ █▀▄   █▄█ █▀█ █▀▄ ██▄ ██▄ ▀▄▀   █ █░▀█   █▄ █▄█ █▄ ▀▀█")
})

var GetTokensHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sl.Info("Запрос на получение токена")

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Host", "localhost")

	var guid Guid
	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(body, &guid); err != nil {
		errHandler(err, "Ошибка при разборе json", &w)
		return
	}

	SendTokenResponse(guid.Guid, &w, db.InsertRefreshToken)

	sl.Info("Токен успешно сгенерирован")
})

var RefreshTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	sl.Info("Запрос на обновление токенов")

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Host", "localhost")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		errHandler(err, "Ошибка при чтении тела запроса", &w)
	}

	token, err := jt.DecodingJsonToken(body)

	if err != nil {
		errHandler(err, "Ошибка при разборе json", &w)
		return
	}

	if token.Refresh == "" || token.Access == "" {
		errHandler(nil, "Отсутствует токен(ы)", &w)
		return
	}

	if claims, err := jwtT.ParseVerifiedAccessToken(token.Access); claims == nil || err != nil {
		errHandler(err, "Ошибка валидации access токена", &w)
		return

	} else {
		if err := jwtT.RefreshTokenValidate(claims.Guid, token.Refresh); err == nil {
			SendTokenResponse(claims.Guid, &w, db.UpdateRefreshToken)

			sl.Info("Токены успешно обновлены")

		} else {
			errHandler(err, "Ошибка валидации refresh токена", &w)
		}
	}
})

func SendTokenResponse(guid string, w *http.ResponseWriter, query func(string, string) error) {

	if guid == "" {
		errHandler(nil, "Поле guid пустое или отсутствует", w)
		return
	}

	access, err := jwtT.GetNewAccessToken(guid)
	if err != nil {
		errHandler(err, "Ошибка при генерации access токена", w)
		return
	}

	refresh, err := jwtT.CreateRefreshToken(guid, query)
	if err != nil {
		errHandler(err, "Ошибка при создании refresh токена", w)
		return
	}

	response, err := jt.TokenEncodingJson(jt.Tokens{Status: 1, Access: access, Refresh: refresh, Guid: guid})
	if err != nil {
		errHandler(err, "Ошибка при кодировании json", w)
		return
	}

	(*w).WriteHeader(http.StatusCreated)

	_, err = (*w).Write(response)
	if err != nil {
		errHandler(err, "Ошибка при записи ответа", w)
		return
	}

	sl.Info("Токен успешно сгенерирован")
}

func errHandler(err error, errText string, w *http.ResponseWriter) {

	sl.ErrorLog(err)

	(*w).WriteHeader(http.StatusBadRequest)

	message, _ := json.Marshal(MsgJson{Status: 0, Msg: errText})
	_, _ = (*w).Write(message)
}
