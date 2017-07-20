# A Tiny golang Docker container example for resin.io

This example demonstrates the use of [Docker multi-stage builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/) to make very small runtime containers. In this example deploy a simple hello world webserver that listens on port 80. The resulting docker container is only `5.26 MB`, compared to `642 MB` of full `resin/raspberrypi3-golang:1.8` golang base image that is a massive improvement. 

After pushing this project to your resin.io device, you should see the following in your dashboard logs:
```
0.07.17 11:31:40 (+0100) Starting application 'registry2.resin.io/wificonnect/2980081f163d08555cc19df54c0fe90708e7bb2f'
20.07.17 11:31:41 (+0100) Started server on port :80.
20.07.17 11:31:41 (+0100) Open https://56b40a6e13416c3f25475a9ed17762da.resindevice.io/ in your browser.
```
**Note:** You will need to enable the [deviceURL](https://docs.resin.io/management/devices/#enable-public-device-url) for your device before opening the link printed in the logs.