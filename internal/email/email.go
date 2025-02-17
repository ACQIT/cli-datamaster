package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
)

func getToken(tenantId, clienteId, clientScret string) string {

	var token Response

	body := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clienteId},
		"client_secret": {clientScret},
		"resource":      {"https://graph.microsoft.com"},
	}

	login := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", tenantId)

	res, err := http.PostForm(login, body)

	defer res.Body.Close()

	if err != nil {
		log.Printf("error al hacer la solicitu %v", err)
		return ""
	}

	if res.StatusCode == 200 {
		data, err := io.ReadAll(res.Body)

		if err != nil {
			log.Printf("error al leer el archivo %v", err.Error())
		}
		err = json.Unmarshal(data, &token)

		if err != nil {
			log.Printf("error al parsear el archivo %v", err.Error())
		}

		return token.AccessToken

	}
	return token.AccessToken
}

func SendEmail(data TempaleteData, userId, ToUser, tenantId, clienteId, clientScret string) {
	var dataEmail EmailRequest
	token := getToken(tenantId, clienteId, clientScret)

	if token == "" {
		log.Println("error al obtener el token")
		return
	}

	htmlContent := generateBodyHtml(data)
	if htmlContent == "" {
		log.Println("error al generar el cuerpo del email")
		return
	}

	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/sendMail", userId)

	dataEmail.Message.Subject = " Información de Alycs"
	dataEmail.Message.Body.ContentType = "HTML"
	dataEmail.Message.Body.Content = htmlContent
	dataEmail.Message.ToRecipients = append(dataEmail.Message.ToRecipients, Recipient{
		EmailAddress: Address{
			Address: ToUser,
		},
	})

	body, err := json.Marshal(dataEmail)
	if err != nil {
		log.Printf(" error a parser de los datos de cuerpo de email : %v", err.Error())
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))

	req.ContentLength = int64(len(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Printf("error al cargar la solicitud  %v ", err)
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		log.Printf(" codigo de peticion a la api de microsoft %d ", res.StatusCode)
		log.Println(" se envio el mensaje ")
	}

	if res.StatusCode >= 400 {
		log.Printf(" codigo de peticion a la api de microsoft %d ", res.StatusCode)
		log.Println(" el mensaje no se pudo enviar el mensaje ")
	}
}

func generateBodyHtml(data TempaleteData) string {
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Printf("error al generar el cuerpo del email %v", err)
		return ""
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, data)

	if err != nil {
		log.Printf("error al generar el cuerpo del email %v", err)
		return ""
	}

	return buff.String()

}
