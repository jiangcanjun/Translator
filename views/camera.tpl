
<html>
<meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no"/>
    <head>
    </head>
	<script>
	function onTake(){	
		var file = document.getElementById("btn_camera").files[0];  
		var reader = new FileReader();
		reader.οnlοad=function(e) {  
		document.getElementById("image").src = e.target.result;	
		}
		reader.readAsDataURL(file);
		
	}
	</script>
    <body>
	<p>open camer demo</p>
	<form id="take_picture">
        <input id="btn_camera" type="file" accept="image/*" capture="camera" οnchange="onTake()" />
	</form>
		<img id="image" width="300" height="200" />
    </body>
</html>