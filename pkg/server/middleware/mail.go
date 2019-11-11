package middleware

import (
	"backend_api/pkg/server/response"
	"encoding/json"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
)

type Mail struct {
	Token string `json:"token"`
	Text  string `json:"text"`
}

func (ctrl *AuthenticateCtrl) SendMail() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// RequestBodyのパース
		var reqBody Mail
		err := json.NewDecoder(request.Body).Decode(&reqBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}

		user, err := ctrl.UserSrv.Repo.SelectByAuthToken(reqBody.Token)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}
		log.Println(user)

		from := mail.NewEmail("OthloQuest運営", "othlotech@gmail.com")
		subject := "プロジェクトへの参加が確定しました！"
		to := mail.NewEmail(user.Name+"様", user.Email)
		plainTextContent := reqBody.Text
		htmlContent := "<strong>" + reqBody.Text + "</strong>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient("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")
		result, err := client.Send(message)
		if err != nil {
			log.Println(err)
		}

		log.Println("result:%s", result)

		//from := mail.NewEmail("Example User", reqBody.From)
		//subject := reqBody.Subject
		//to := mail.NewEmail("Example User", reqBody.To)
		//plainTextContent := reqBody.PlainText
		//htmlContent := reqBody.HtmlContent
		//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		//client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		//_, err = client.Send(message)
		//if err != nil {
		//	log.Println(err)
		//}
	}
}
