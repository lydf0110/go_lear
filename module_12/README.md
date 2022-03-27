# go_lear
#本地验证，没加证书，
#openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:2048 -subj '/O=test-home Inc./CN=*.test' -keyout test.key -out test.crt
#kubectl create -n istio-system secret tls wildcard-credential --key=test.cafe.key --cert=test.cafe.crt
#整体过程，步骤说明下；
#创建gateway
#创建vs
#然后通过刷流量，检测拓扑图监控