<html>
<head>
    <title> ファイルアップロード </title>
</head>
<body>
    <form enctype="multipart/form-data" action="/form" method="post">
        <input type="file" name="uploadfile" />
        <input type="hidden" name="token" value="{{.}}"/>
        <input type="submit" value="upload"/>
    </form>
</body>
</html>