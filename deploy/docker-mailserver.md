# How to setup docker-mailserver

1. Create the file docker-compose.yml with the following content:

    ```yml
    services:
        docker-mailserver1:
            image: ghcr.io/docker-mailserver/docker-mailserver:latest
            hostname: mail.example.com
            ports:
            - "25:25"
            - "587:587"
            - "465:465"
            volumes:
            - ./docker-data/dms/mail-data/:/var/mail/
            - ./docker-data/dms/mail-state/:/var/mail-state/
            - ./docker-data/dms/mail-logs/:/var/log/mail/
            - ./docker-data/dms/config/:/tmp/docker-mailserver/
            - /etc/localtime:/etc/localtime:ro
            - /etc/letsencrypt:/etc/letsencrypt
            environment:
            - ENABLE_RSPAMD=1
            - ENABLE_CLAMAV=1
            - ENABLE_FAIL2BAN=1
            - SSL_TYPE=letsencrypt
            - PERMIT_DOCKER=network
            - SPOOF_PROTECTION=0
            cap_add:
            - NET_ADMIN # For Fail2Ban to work
            restart: always
    ```

    Replace `mail.example.com` with the FQDN of your mail server.

2. Start docker-mailserver

    ```
    docker compose up -d
    ```

3. Configure your DNS service
    - set an `MX` record for your domain. The `MX` record contains the name of your mail server (for example `mail.example.com`)
    - set an `A` record that resolves the name of your mail server. The `A` record contains the IP address that your mail server resoles to (for example `11.22.33.44`)
    - set a `PTR` record that resolves the IP of your mail server

    You can use `dig` to check if you configured everything properly:
    ```bash
    $ dig @1.1.1.1 +short MX example.com
    mail.example.com
    $ dig @1.1.1.1 +short A mail.example.com
    11.22.33.44
    $ dig @1.1.1.1 +short -x 11.22.33.44
    mail.example.com
    ```

4. Setup DKIM
    - Generate the DKIM files
    ```
    docker exec -ti <CONTAINER_NAME> setup config dkim
    ```
    - set a `TXT` record for `mail._domainkey` that contains the value generated from the previous step

5. Setup DMARC
    - set a `TXT` record for `dmarc` that contains the following value:
        ```
        "v=DMARC1; p=none; sp=none; fo=0; adkim=4; aspf=r; pct=100; rf=afrf; ri=86400; rua=mailto:dmarc.report@example.com; ruf=mailto:dmarc.report@example.com"
        ```
6. Setup SPF
    - set a `TXT` record for `<hostname>` that contains the following value:
        ```
        v=spf1 mx ~all
        ```

7. Create email account to use docker-mailserver
    ```
    docker exec -ti <CONTAINER_NAME> setup email add <EMAIL> <PASSWORD>
    ```

8. Now you can send email using docker-mailserver! Here's a simple example code to send an email:
    ```go
    package main

    import (
        "github.com/wneessen/go-mail"
        "log"
    )

    func main() {
        // Create an email message
        m := mail.NewMsg()
        if err := m.From("sender@testmail.benalu.dev"); err != nil {
            log.Fatalf("failed to set From address: %s", err)
        }
        if err := m.To("receiver@testmail.benalu.dev"); err != nil {
            log.Fatalf("failed to set To address: %s", err)
        }
        m.Subject("Email Subject")
        m.SetBodyString(mail.TypeTextPlain, "Email Content")

        // Setup the email client
        c, err := mail.NewClient("benalu.dev",
            mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
            mail.WithUsername("admin@testmail.benalu.dev"), mail.WithPassword("password"),
            mail.WithTLSPolicy(mail.TLSMandatory))
        if err != nil {
            log.Fatalf("failed to create mail client: %s", err)
        }

        // Send out the email
        if err := c.DialAndSend(m); err != nil {
            log.Fatalf("failed to send mail: %s", err)
        }
    }
    ```