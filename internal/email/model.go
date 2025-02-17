package email

type EmailRequest struct {
	Message Message `json:"message"`
}

type Message struct {
	Subject      string      `json:"subject"`
	Body         Body        `json:"body"`
	ToRecipients []Recipient `json:"toRecipients"`
}

type Body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type Recipient struct {
	EmailAddress Address `json:"emailAddress"`
}

type Address struct {
	Address string `json:"address"`
}

type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Resource    string `json:"resource"`
}

const emailTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nueva ALYC Disponible</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333333;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        .header {
            background-color: #0078D4;
            color: white;
            padding: 20px;
            text-align: center;
        }
        .content {
            padding: 20px;
            background-color: #ffffff;
        }
        .footer {
            background-color: #f5f5f5;
            padding: 20px;
            text-align: center;
            font-size: 12px;
        }
        .email-button {
            font-family: inherit;
            font-size: 20px;
            background: royalblue;
            color: white;
            padding: 0.7em 1em;
            padding-left: 0.9em;
            display: inline-flex;
            align-items: center;
            border: none;
            border-radius: 16px;
            overflow: hidden;
            text-decoration: none;
            margin: 20px 0;
        }
        .svg-wrapper-1 {
            display: flex;
            align-items: center;
            margin-right: 0.3em;
        }
        .email-button svg {
            display: block;
            width: 24px;
            height: 24px;
        }
        .details {
            background-color: #f9f9f9;
            padding: 15px;
            border-radius: 4px;
            margin: 15px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Nueva ALYC Disponible</h1>
        </div>
        <div class="content">
            <p>Estimado/a {{.NombreCliente}},</p>
            
            <p>Nos complace informarle que hay una nueva ALYC disponible .</p>
            
            <div class="details">
                <h3>Detalles de la ALYC</h3>
                <p><strong>Mercado :</strong> {{.Market}}</p>
                <p><strong>Nombre:</strong> {{.NombreALYC}}</p>
                <p><strong>Telefono:</strong> {{.Phone}}</p>
                <p><strong>Email:</strong> {{.EmailALYC}}</p>
            </div>
            <p>Para ver más detalles y acceder a la información completa, haga clic en el siguiente botón:</p>
            
            <a href="{{.URLDetalle}}" class="email-button">
                <div class="svg-wrapper-1">
                    <div class="svg-wrapper">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                            <path fill="none" d="M0 0h24v24H0z"></path>
                            <path
                                fill="currentColor"
                                d="M1.946 9.315c-.522-.174-.527-.455.01-.634l19.087-6.362c.529-.176.832.12.684.638l-5.454 19.086c-.15.529-.455.547-.679.045L12 14l6-8-8 6-8.054-2.685z"
                            ></path>
                        </svg>
                    </div>
                </div>
                <span>Ver Detalles</span>
            </a>
            
            <p>Si tiene alguna pregunta o necesita asistencia adicional, no dude en contactarnos.</p>
            
            <p>Saludos cordiales,<br>
            El equipo de {{.NombreEmpresa}}</p>
        </div>
        <div class="footer">
            <p>Este es un correo automático. Por favor, no responda a este mensaje.</p>
            <p>{{.NombreEmpresa}} - Todos los derechos reservados </p>
        </div>
    </div>
</body>
</html>`

type TempaleteData struct {
	Market        string
	NombreCliente string
	NombreALYC    string
	URLDetalle    string
	NombreEmpresa string
	EmailALYC     string
	Phone         string
}
