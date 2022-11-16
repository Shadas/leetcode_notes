https://toutyrater.github.io/prep/start.html

config.json

{
    "log": {
        "access": "", 
        "error": "", 
        "loglevel": "info"
    }, 
    "inbounds": [
        {
            "port": 80, // port 
            "protocol": "vmess", 
            "settings": {
                "clients": [
                    {
                        "id": "", // uuid 
                        "alterId": 0, 
                        "security": "aes-128-gcm"
                    }
                ]
            }
        }
    ], 
    "outbounds": [
        {
            "protocol": "freedom", 
            "settings": {
                "response": null
            }, 
            "tag": "direct"
        }, 
        {
            "protocol": "blackhole", 
            "settings": {
                "response": {
                    "type": "http"
                }
            }, 
            "tag": "cnblock"
        }
    ], 
    "dns": {
        "servers": [
            "8.8.8.8", 
            "8.8.4.4", 
            "localhost"
        ]
    }, 
    "routing": {
        "domainStrategy": "IPIfNonMatch", 
        "rules": [
            {
                "type": "field", 
                "ip": [
                    "geoip:private"
                ], 
                "outboundTag": "direct"
            }, 
            {
                "type": "field", 
                "domain": [
                    "geosite:cn"
                ], 
                "outboundTag": "direct"
            }, 
            {
                "type": "field", 
                "domain": [
                    "geoip:cn"
                ], 
                "outboundTag": "direct"
            }, 
            {
                "type": "field", 
                "ip": [
                    "geoip:cn", 
                    "geoip:private"
                ], 
                "outboundTag": "cnblock"
            }
        ]
    }
}
