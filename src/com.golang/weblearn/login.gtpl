<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post" >
    Username:<input type="text" name="username">
    Password:<input type="password" name="password">
    Verification code:<input type="verificationCode" name="verificationCode">
    <input type="submit" value="Login">
</form>
<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" name="uploadfile" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>
</body>
</html>
