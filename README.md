# JsValidator
JsValidator is a tool created for validating the JS files after crawlling it from waybackurls

# Idea
When you get js files from waybackruls, some of the urls is bad and others are 404, this script is filtering all of them and outting only the valid js files directly <3

# Installation
Just need to install go, run:

```
▶ brew install go
▶ git clone https://github.com/SirBugs/JsValidator.git
```

or download from https://go.dev/dl/

# Usage:
```
▶ go run main.go urls.txt

    /$$$$$           /$$    /$$          /$$ /$$       /$$             /$$
   |__  $$          | $$   | $$         | $$|__/      | $$            | $$
      | $$  /$$$$$$$| $$   | $$ /$$$$$$ | $$ /$$  /$$$$$$$  /$$$$$$  /$$$$$$    /$$$$$$   /$$$$$$
      | $$ /$$_____/|  $$ / $$/|____  $$| $$| $$ /$$__  $$ |____  $$|_  $$_/   /$$__  $$ /$$__  $$
 /$$  | $$|  $$$$$$  \  $$ $$/  /$$$$$$$| $$| $$| $$  | $$  /$$$$$$$  | $$    | $$  \ $$| $$  \__/
| $$  | $$ \____  $$  \  $$$/  /$$__  $$| $$| $$| $$  | $$ /$$__  $$  | $$ /$$| $$  | $$| $$
|  $$$$$$/ /$$$$$$$/   \  $/  |  $$$$$$$| $$| $$|  $$$$$$$|  $$$$$$$  |  $$$$/|  $$$$$$/| $$
 \______/ |_______/     \_/    \_______/|__/|__/ \_______/ \_______/   \___/   \______/ |__/

                             JsValidator Tool By @SirBugs .go Version
                                    V: 1.0.1 Made With All Love
                        For Validating Javascript files out of waybackurls
                                Twitter@SirBagoza / GitHub@SirBugs
                                     Run: go run main.go file
[ 200 ] :: https://academy.launchdarkly.com/static/js/jquery.min.be3fe6ef3675.js
[ 200 ] :: https://academy.launchdarkly.com/static/js/modernizr.min.c89684367713.js
[ 200 ] :: https://academy.launchdarkly.com/static/js/plugins.min.a45d9e4c0c5a.js
[ 404 ] :: https://academy.launchdarkly.com/static/js/api/api-service.e4a79943a76c.js
[ 404 ] :: https://academy.launchdarkly.com/static/js/catalog.min.29ad8fe27fb1.js
[ 200 ] :: https://academy.launchdarkly.com/static/js/vendor/axios.min.13f25a468bb3.js
[ 200 ] :: https://academy.launchdarkly.com/static/js/vendor/prism/prism.d722a89f1d58.js
[ 404 ] :: https://academy.launchdarkly.com/static/js/sanitize-html.931242c1ddfb.js
[ 404 ] :: https://academy.launchdarkly.com/static/js/scripts.min.59eb9224163a.js
[ 404 ] :: https://apidocs.launchdarkly.com/codemirror/mode/javascript/javascript.js
[ 404 ] :: https://apidocs.launchdarkly.com/codemirror/mode/http/http.js
[ 404 ] :: https://apidocs.launchdarkly.com/codemirror/mode/shell/shell.js
[ 404 ] :: http://apidocs.launchdarkly.com/angular-marked.js
[ 404 ] :: https://apidocs.launchdarkly.com/2.4.0/redocly-reference-docs.min.js

```

# One Line Command:

```
▶ echo 'target.com' | waybackurls | tee waybackresults.txt; cat waybackresults.txt | grep "\.js" > js_files.txt; go run main.go js_files.txt
```

// You can use Gau, HaKrawler, Katana, etc...

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3
Twitter@SirBagoza , Github@SirBugs
