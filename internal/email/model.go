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
        .button {
            display: inline-block;
            padding: 10px 20px;
            background-color: #0078D4;
            color: white;
            text-decoration: none;
            border-radius: 4px;
            margin: 20px 0;
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
            
            <p>Nos complace informarle que hay una nueva ALYC disponible en nuestro sistema.</p>
            
            <div class="details">
                <h3>Detalles de la ALYC:</h3>
                <p><strong>Nombre:</strong> {{.NombreALYC}}</p>
            </div>

            <p>Para ver más detalles y acceder a la información completa, haga clic en el siguiente botón:</p>
            
            <a href="{{.URLDetalle}}" class="button">Ver Detalles</a>
            
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
	NombreCliente string
	NombreALYC    string
	URLDetalle    string
	NombreEmpresa string
}
