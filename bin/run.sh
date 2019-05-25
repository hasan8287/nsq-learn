

#SETUP
export NSQD_HOST=127.0.0.1:4150
export NSQD_HOST_CONSUME=127.0.0.1:4160

if [[ ${1} == 'publish' || ${1} == ''  ]]; then
	go run "./publish/main.go"
fi

if [[ ${1} == 'consume' || ${1} == ''  ]]; then
	go run "./consume/main.go"
fi