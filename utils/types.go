package utils

type Response struct {
	Message string `json:"message"`
}

type PlaygroundTemplateData struct {
	Endpoint string
}

// Credits: Ice <3
const PlaygroundTemplate = `
{{ define "index" }}
<!DOCTYPE html>
<html>
    <head>
        <title>Nino - GraphQL Playground</title>
        <meta charset="UTF-8" />
        <link rel="shortcut icon" href="https://cdn.floofy.dev/images/trans.png" />
        <link rel="icon" href="https://cdn.floofy.dev/images/trans.png" />
        <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui" />
        <link href="//fonts.googleapis.com/css?family=Open+Sans:300,400,600,700|Source+Code+Pro:400,700" rel="stylesheet" />
        <link href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css" rel="stylesheet" />
        <script src="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
    </head>

    <body>
        <div id="root">
			<style>
			  body {
				background-color: rgb(23, 42, 58);
				font-family: Open Sans, sans-serif;
				height: 90vh;
			  }
		
			  #root {
				height: 100%;
				width: 100%;
				display: flex;
				align-items: center;
				justify-content: center;
			  }
		
			  .loading {
				font-size: 32px;
				font-weight: 200;
				color: rgba(255, 255, 255, .6);
				margin-left: 20px;
			  }
		
			  img {
				width: 78px;
				height: 78px;
			  }
		
			  .title {
				font-weight: 400;
			  }
			</style>

			<img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt='graphql logo' />
			<div class="loading">
				<span class="title">GraphQL Playground</span>
			</div>
		</div>
        <script>
			document.addEventListener('load', () => GraphQLPlayground.init(document.getElementById('root'), {
				endpoint: {{ .Endpoint }},
				'general.betaUpdates': true,
				'tracing.hideTracingResponse': false,
				'tracing.tracingSupported': false
			}));
		</script>
    </body>
</html>
{{ end }}
`
