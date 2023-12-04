/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "footer"}}
  </div><!-- .container -->	
	<div class="container">
      <div class="text-footer">
        <div >&copy; 2019 {{.HomePage}} qq:1269866868 mail:golangbbs@gmail.com  <img src="/static/img/about.png" alt="" style="vertical-align: middle;margin-left:8px;margin-bottom:8px;width: 20px;height:20px;cursor: pointer;" onclick="location.href='/about'"></div>
		</div>
    </div>
		{{if .Csrf}}<script type="text/javascript">window.csrf_token="{{.Csrf}}";</script>{{end}}
  </body>
</html>
<script src="/static/footer.js"></script>
<script src="/static/js/cat/L2Dwidget.min.js"></script>
<script type="text/javascript">
try{
    L2Dwidget.init({
        "model": {
            jsonPath: "/static/js/cat/tororo.model.json",
            "scale": 1
        },
        "display": {
            "position": "left",
            "width": 150,
            "height": 300,
            "hOffset": -70,
            "vOffset": -310
        },
        "mobile": {
            "show": true,
            "scale": 0.5
        },
        "react": {
            "opacityDefault": 0.7,
            "opacityOnHover": 0.2
        }
    });
}catch(e){}
</script>
<div id="live2d-widget"><canvas id="live2dcanvas" width="300" height="600" style="position: fixed; opacity: 0.7; left: 0px; top: 20px; z-index: 9999; pointer-events: none;" class=""></canvas></div>
{{end}}
