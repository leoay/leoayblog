<!DOCTYPE html>

<html>
<head>
  <title>leoay lab</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
</body>

<script>
    var ua = navigator.userAgent.toLowerCase();
    var isWeixin = ua.indexOf('micromessenger') != -1;
    if (!isWeixin) {
      window.location.href = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=0"
    }else {
        
    }
</script>