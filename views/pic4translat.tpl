<html>
	<head>
		<title>上传文件</title>
	</head>
	<body>
		<form enctype="multipart/form-data" action="/pic4trans" method="post">
			<input type="file" name="uploadfile">
			<input type="hidden" name="token" value="{...{.}...}">
			<input type="submit" value="translat">
		</form>
	</body>
</html>
