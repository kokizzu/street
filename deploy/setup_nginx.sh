
sudo apt install nginx certbot python3-certbot-nginx
echo '
server {
  listen [::]:443 ssl; # managed by Certbot
  listen 443 ssl; # managed by Certbot
  ssl_certificate /etc/letsencrypt/live/admin.hapstr.xyz/fullchain.pem; # managed by Certbot
  ssl_certificate_key /etc/letsencrypt/live/admin.hapstr.xyz/privkey.pem; # managed by Certbot
  include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

  server_name admin.hapstr.xyz;

  location / {
    root /home/street/svelte;
    try_files $uri @apiproxy;
  }

  location @apiproxy {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarwwwwwwwwwwwwded-Proto $scheme;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:1234;
  }
}
' | sudo tee /etc/nginx/sites-available/admin.hapstr.xyz
sudo ln -s /etc/nginx/sites-available/admin.hapstr.xyz /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
sudo certbot --nginx -d admin.hapstr.xyz -d admin.hapstr.xyz
sudo systemctl enable certbot.timer
