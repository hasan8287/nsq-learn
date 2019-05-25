

#SETUP
if [[ ${1} == 'setup' ]]; then
	export NSQD_HOST=127.0.0.1:4150
fi

if [[ ${1} == 'publish' || ${1} == ''  ]]; then
	go run "./publish/main.go"
fi