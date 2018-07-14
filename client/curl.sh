curl -i -N \
-H "Connection: Upgrade" \
-H "Upgrade: websocket" \
-H "Host: echo.websocket.org" \
-H "Origin: https://www.websocket.org" \
http://localhost:3000/ws