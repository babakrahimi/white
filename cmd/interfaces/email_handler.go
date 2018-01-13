package interfaces

import (
	"bytes"
	"html/template"

	"github.com/megaminx/white/cmd/business/user"
)

type (
	EmailServerProvider interface {
		Send(to, subject, content string) error
	}

	EmailHandler struct {
		Provider                EmailServerProvider
		RegistrationRedirectURL string
	}
)

func NewEmailHandler(p EmailServerProvider, redirectURL string) user.EmailHandler {
	eh := &EmailHandler{
		Provider:                p,
		RegistrationRedirectURL: redirectURL,
	}
	return eh
}
func (eh *EmailHandler) SendInvitationEmail(to, token string) error {
	subject := "عضویت در سامانه"
	content, err := GetInvitationEmailContent(eh.RegistrationRedirectURL, token)
	if err != nil {
		return err
	}

	if err := eh.Provider.Send(to, subject, content); err != nil {
		return err
	}
	return nil
}

func GetInvitationEmailContent(url, token string) (string, error) {
	body := `
	    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
      <title>Mahan Team</title>
      <style type="text/css">
      body {
       padding-top: 0 !important;
       padding-bottom: 0 !important;
       padding-top: 0 !important;
       padding-bottom: 0 !important;
       margin:0 !important;
       width: 100% !important;
       -webkit-text-size-adjust: 100% !important;
       -ms-text-size-adjust: 100% !important;
       -webkit-font-smoothing: antialiased !important;
     }
     .tableContent img {
       border: 0 !important;
       display: block !important;
       outline: none !important;
     }
     a{
      color:#382F2E;
    }

    p, h1{
      color:#382F2E;
      margin:0;
    }
 p{
      text-align:right;
      color:#999999;
      font-size:12px;
      font-weight:normal;
      line-height:19px;
      direction: rtl;
    }

    a.link1{
      color:#382F2E;
    }
    a.link2{
      font-size:12px;
      font-family: tahoma, arial, sans-serif;
      text-decoration:none;
      color:#ffffff;
    }

    h2{
      text-align:right;
       color:#222222; 
       font-size:12px;
      font-weight:bold;
      direction: rtl;
    }
    div,p,ul,h1{
      margin:0;
    }

    .bgBody{
      background: #ffffff;
    }
    .bgItem{
      background: #ffffff;
    }
	
@media only screen and (max-width:480px)
		
{
		
table[class="MainContainer"], td[class="cell"] 
	{
		width: 100% !important;
		height:auto !important; 
	}
td[class="specbundle"] 
	{
		width:100% !important;
		float:left !important;
		font-size:13px !important;
		line-height:17px !important;
		display:block !important;
		padding-bottom:15px !important;
	}
		
td[class="spechide"] 
	{
		display:none !important;
	}
	    img[class="banner"] 
	{
	          width: 100% !important;
	          height: auto !important;
	}
		td[class="left_pad"] 
	{
			padding-left:15px !important;
			padding-right:15px !important;
	}
		 
}
	
@media only screen and (max-width:540px) 

{
		
table[class="MainContainer"], td[class="cell"] 
	{
		width: 100% !important;
		height:auto !important; 
	}
td[class="specbundle"] 
	{
		width:100% !important;
		float:left !important;
		font-size:13px !important;
		line-height:17px !important;
		display:block !important;
		padding-bottom:15px !important;
	}
		
td[class="spechide"] 
	{
		display:none !important;
	}
	    img[class="banner"] 
	{
	          width: 100% !important;
	          height: auto !important;
	}
	.font {
		font-size:18px !important;
		line-height:22px !important;
		
		}
		.font1 {
		font-size:18px !important;
		line-height:22px !important;
		
		}
}

    </style>
<script type="colorScheme" class="swatch active">
{
    "name":"Default",
    "bgBody":"ffffff",
    "link":"382F2E",
    "color":"999999",
    "bgItem":"ffffff",
    "title":"222222"
}
</script>
  </head>
  <body paddingwidth="0" paddingheight="0"   style="padding-top: 0; padding-bottom: 0; padding-top: 0; padding-bottom: 0; background-repeat: repeat; width: 100% !important; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; -webkit-font-smoothing: antialiased;" offset="0" toppadding="0" leftpadding="0">
    <table bgcolor="#ffffff" width="100%" border="0" cellspacing="0" cellpadding="0" class="tableContent" align="center"  style='font-family:Tahoma, Helvetica, Arial,serif;'>
  <tbody>
    <tr>
      <td><table width="600" border="0" cellspacing="0" cellpadding="0" align="center" bgcolor="#ffffff" class="MainContainer">
  <tbody>
    <tr>
      <td><table width="100%" border="0" cellspacing="0" cellpadding="0">
  <tbody>
    <tr>
      <td valign="top" width="40">&nbsp;</td>
      <td><table width="100%" border="0" cellspacing="0" cellpadding="0">
  <tbody>
  <!-- =============================== Header ====================================== -->   
    <tr>
    	<td height='75' class="spechide"></td>
        
        <!-- =============================== Body ====================================== -->
    </tr>
    <tr>
      <td class='movableContentContainer ' valign='top'>

        <div class="movableContent" style="border: 0px; padding-top: 0px; position: relative;">
        	<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
                          <tr>
                            <td align='left'>
                              <div class="contentEditableContainer contentTextEditable">
                                <div class="contentEditable" align='center'>
                                  <h2 >سلام همکار عزیز،</h2>
                                </div>
                              </div>
                            </td>
                          </tr>

                          <tr><td height='15'> </td></tr>

                          <tr>
                            <td align='left'>
                              <div class="contentEditableContainer contentTextEditable">
                                <div class="contentEditable" align='center'>
                                  <p>
                                    لطفا برای عضویت و تکمیل مشخصات فردی خود در سامانه تیم نرم‌افزار ماهان، بر روی لینک زیر کلیک کنید.
                                  </p>
                                </div>
                              </div>
                            </td>
                          </tr>

                          <tr><td height='25'></td></tr>

                          <tr>
                            <td align='center'>
                              <table>
                                <tr>
                                  <td align='center' bgcolor='#1A54BA' style='background:#69C374; padding:15px 18px;-webkit-border-radius: 4px; -moz-border-radius: 4px; border-radius: 4px;'>
                                    <div class="contentEditableContainer contentTextEditable">
                                      <div class="contentEditable" align='center'>
                                        <a target='_blank' href='{{.URL}}?t={{.Token}}' class='link2' style='color:#ffffff;'>{{.URL}}</a>
                                      </div>
                                    </div>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                          <tr>
                            <td align='left'>
                              <div class="contentEditableContainer contentTextEditable">
                                <div class="contentEditable" align='center'>
                                  <p>با تشکر</p>
                                  <p>سامانه تیم نرم‌افزار ماهان</p>
                                </div>
                              </div>
                            </td>
                          </tr>
                          <tr><td height='20'></td></tr>
                        </table>
        </div>
        <div class="movableContent" style="border: 0px; padding-top: 0px; position: relative;">
        	<table width="100%" border="0" cellspacing="0" cellpadding="0">
  <tbody>
    <tr>
      <td height='65'>
    </tr>

    <tr><td height='25'></td></tr>
    <tr>
      <td><table width="100%" border="0" cellspacing="0" cellpadding="0">

</table>
</td>
    </tr>
    <tr><td height='88'></td></tr>
  </tbody>
</table>

        </div>
        
        <!-- =============================== footer ====================================== -->
      
      </td>
    </tr>
  </tbody>
</table>
</td>
      <td valign="top" width="40">&nbsp;</td>
    </tr>
  </tbody>
</table>
</td>
    </tr>
  </tbody>
</table>
</td>
    </tr>
  </tbody>
</table>


      </body>
      </html>
  
`
	t := template.Must(template.New("email").Parse(body))
	b := bytes.NewBufferString("")
	err := t.Execute(b, struct {
		URL, Token string
	}{
		URL:   url,
		Token: token,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
