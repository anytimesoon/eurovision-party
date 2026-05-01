# Eurovision Party

Ladies and gentlemen, welcome to the Eurovision Party repository! Get ready to immerse yourself in the dazzling world of Eurovision like never before. 
This project is here to open up how we celebrate this iconic song contest by helping you virtually connect with your friends!

## Features

- **Easy invites** – Users can easily invite friends to help adoption
- **Instant messaging** – Users can discuss the show as its happening live
- **Media sharing** – Users can share photos and videos in the chat.
- **Message reactions** – Users can react to messages with emojis or quick reaction options.
- **Gesture-based replies** – Users on mobile can reply to messages using intuitive gestures for a smoother chat experience.
- **Category-based voting** – Each artist can receive votes in four separate voting categories: song, performance, costume and staging

## Getting Started

This project can be deployed using docker compose:

```yaml
services:
  eurovision-party:
    image: anytimesoon/eurovision-party
    container_name: eurovision-party
    ports:
      - "3000:80"
    restart: unless-stopped
    environment:
      - DOMAIN_NAME=my.domain.com
    volumes:
      - /my/local/data/dir/:/backend/
```

Possible environment variables:

| Variable           | Function                                                      | Default   | Required |
|--------------------|---------------------------------------------------------------|-----------|----------|
| DOMAIN_NAME        | domain name used to connect to the application                | localhost | yes      |
| INSECURE           | use http schema                                               | false     | no       |
| VOTE_COUNT_TRIGGER | how many votes trigger a reminder message to vote in the chat | 5         | no       |
| MAX_INVITES        | number of invites friends can send out to their friends       | 5         | no       |


# Attributions
Thank you to openmoji for designing all emojis – the open-source emoji and icon project. Licence: CC BY-SA 4.0