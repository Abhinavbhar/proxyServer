![proxy](https://github.com/user-attachments/assets/cd771541-e21b-4412-afc6-21919e9df99e)
# ProxyServer
A proxy server built using Golang.

## Use Cases

1. **Access Banned Websites on a WiFi Network**
   - Example: Access Netflix not available on your college WiFi.

2. **Access Banned Websites in Your Country**
   - Example: Many websites like Stake are banned in India.

3. **Watch Content from a Different Region**
   - Example: Netflix in India often censors violence and nudity. Using the proxy server, you can access uncensored versions.

## Basic Design
- Users will add the proxy in their network settings (available on all devices).
- Any traffic from the user will first reach the `ProxyServer`, which will then forward it to the actual websites the user is trying to access.
- The user's Internet provider will believe the traffic is destined for our server instead of the original website, allowing the traffic (until the Internet provider or WiFi admin detects and bans our server as a proxy).

> **Note:** There are ways to mitigate the issue of detection and banning, but this will be addressed only at scale, so it is not a priority for now.

### Flow of User Request
```
User -> wifiAdmin/ISP -> ProxyServer -> Netflix -> ProxyServer ->wifiAdmin/ISP-> User
```

## Tech Stack
1. **Backend:** Golang with the Fiber framework.
2. **Database:** Redis and a custom in-memory store.
3. **Server:** Oracle Virtual Machine located in a country that have unrestricted access to internet.

## Authentication
- Authentication and bandwidth restriction are **not priorities** at the moment.
- These features will be added after the main functionality is built.

## Priority
- Focus on making the proxy functionality work .
- Users should be able to access unrestricted websites using our proxy.

> **Frontend:** Not required initially.
> The proxy server's IP will be made available in the GitHub repository, enabling users to configure and use the proxy directly.

---

