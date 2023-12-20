#curl -X GET "https://api.cloudflare.com/client/v4/user/tokens/verify" \
#     -H "Authorization: Bearer AAAPutYourTokenInHereSoYouCanTestItL5Cl3" \
#     -H "Content-Type:application/json"

# https://api.cloudflare.com/client/v4/zones/27b900d9e05cfb9f3a64fecff2497f90/dns_records
#
# {
#    "comment": "WIT DNS Control Panel",
#    "content": "2001:4860:4860::8888",
#    "name": "www",
#    "proxied": false,
#    "ttl": 3600,
#    "type": "AAAA"
#}

curl --request POST \
  --url https://api.cloudflare.com/client/v4/zones/27b900d9e05cfb9f3a64fecff2497f90/dns_records \
  --header 'Content-Type: application/json' \
  --header 'X-Auth-Key: e08806ad85ef97aebaacd2d7fa462a7d417a7x' \
  --header 'X-Auth-Email: basilarchia@gmail.com' \
  --data '{
  "comment": "WIT DNS Control Panel",
  "content": "2001:4860:4860::5555",
  "name": "www5",
  "proxied": false,
  "ttl": 3600,
  "type": "AAAA"
}'
