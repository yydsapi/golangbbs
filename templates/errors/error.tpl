/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "showerrinfo"}}
<!DOCTYPE html>
<html>
<head>
	<title>Error Page</title>
	<link rel="stylesheet" href="/static/error.css">
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
	<div class="main">
				<div class="text">
					<p><a href="/" title="return"><img src="/static/img/return2.png" class="imgreturn"></a> tips:</p>
					<p></p><p></p><p></p>
					<h2>{{.info}}</h2>	
					<h2>{{.else}}</h2>
					<h4><a href=/signin><font color=white >click to login</font></a></h4>
					<p></p>
				</div>
				<div class="image">
					<img src="/static/img/smile.png">
				</div>
				<div class="clear"></div>
			<div class="footer">
				<p>&copy; 2019 golangbbs.com qq:1269866868 mail:golangbbs@gmail.com <img src="/static/img/about.png" alt="" style="vertical-align: middle;margin-left:8px;margin-bottom:8px;width: 20px;height:20px;cursor: pointer;" onclick="location.href='/about'"></p>
			</div>
	</div>

</body>
</html>
{{end}}
