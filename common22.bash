#suomiporno.xyz OK
#bin/newsite -locale=fi_FI -themes=porno -site=suomiporno.xyz -cssthemes=stou-dk-theme
#bin/newsite -locale=fi_FI -themes=porno -site=www.suomiporno.xyz -cssthemes=stou-dk-theme
#bin/newsite -locale=fi_FI -themes=porno -site=blog.suomiporno.xyz  -cssthemes=stou-dk-theme

bin/jsontomarkdown -locale=fi_FI -themes=porno -site=suomiporno.xyz
bin/jsontomarkdown -locale=fi_FI -themes=porno -site=www.suomiporno.xyz
bin/jsontomarkdown -locale=fi_FI -themes=porno -site=blog.suomiporno.xyz

hugo -s www/fi_FI/porno/suomiporno.xyz
hugo -s www/fi_FI/porno/www.suomiporno.xyz
hugo -s www/fi_FI/porno/blog.suomiporno.xyz
bin/sitemapsping -locale=fi_FI -themes=porno -site=suomiporno.xyz
bin/sitemapsping -locale=fi_FI -themes=porno -site=www.suomiporno.xyz
bin/sitemapsping -locale=fi_FI -themes=porno -site=blog.suomiporno.xyz
