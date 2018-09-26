<!DOCTYPE html>
<html lang="en">
{{ .Header }}
<body>
<div class="container-fluid">
    {{ .Nav }}
    {{ template "components/footer.tpl" }}
    {{ .LayoutContent }}
</div>
{{ compress_js "lib" }}
{{ compress_js "app" }}
{{ .Scripts }}
</body>
</html>
