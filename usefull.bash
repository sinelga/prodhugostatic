hugo server -b http://nainti.club -w --appendPort=false --disableLiveReload

bin/newsite -locale=fi_FI -themes=porno -site=blog.nainti.biz -cssthemes=hyde

bin/newsite -locale=fi_FI -themes=porno -site=www.nainti.biz -cssthemes=hyde
bin/newsite -locale=fi_FI -themes=porno -site=nainti.biz -cssthemes=hyde

bin/jsontomarkdown -site=blog.nainti.biz

hugo -t herring-cove

hugo -t hyde


{{ .Site | printf "%+v" }}


???curl -I -d "sitemap=http://seuraa.co" http://www.google.com/webmaster/tools/ping

find . | grep .git | xargs rm -rf
find . | grep /.git$ | xargs rm -rf #more correct
