<html>
<head>
<title></title>
</head>
<body>
<form action="/form?username=kitoria" method="post">
    ユーザ名:<input type="text" name="username">
    パスワード:<input type="password" name="password">
    <br><br>
    <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banane">banane</option>
    </select>
    <br><br>

    <input type = "radio" name = "gender" value="1">男
    <input type = "radio" name = "gender" value="2">女
    <br><br>
    
    <input type="checkbox" name="interest" value="football">サッカー
    <input type="checkbox" name="interest" value="basketball">バスケットボール
    <input type="checkbox" name="interest" value="tennis">テニス
    <br><br>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="提出">
</form>
</body>
</html>