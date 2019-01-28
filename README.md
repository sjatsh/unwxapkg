[![Build Status](https://travis-ci.org/sjatsh/unwxapkg.svg?branch=master)](https://travis-ci.org/sjatsh/unwxapkg)
[![codecov](https://codecov.io/gh/sjatsh/unwxapkg/branch/master/graph/badge.svg)](https://codecov.io/gh/sjatsh/unwxapkg)

## Description
 
 The tool is used to decompress the wxapkg compress of the Wechat applet. Wxapkg can be obtained through NetEase 
 [MuMu](http://mumu.163.com/) simulator. 
 
 - Download it and install wechat
 ![](http://static.1sapp.com/simage_template/401f78fc5c26cefb839d7c37fb2451fe39364d86.png)
 - Get applet wxapkg package use 'RE File Manager'
 go to `/data/data/com.tencent.mm/MicroMsg/{{a hexadecimal string folder}}/appbrand/pkg/`
 ![](http://static.1sapp.com/simage_template/843604b9abd8c859c3de0eea50ec1a821892dc21.png)
 - Compress and send
 ![](http://static.1sapp.com/simage_template/50a67d45e553480e817ced99797ec17e0029ee33.png)
 ![](http://static.1sapp.com/simage_template/a16e9408afebc1d70430512d066e9fd476cb70ef.png)
 

## Usage

```
go get github.com/sjatsh/unwxapkg
cd ~/go/src/github.com/sjatsh/unwxapkg
unwxapkg -f dest/102.wxapkg
```

## Unpacking Result
![](http://static.1sapp.com/simage_template/23fa85f16911f689d7f35de36c9fd725bac75549.png)

## Can be use by
微信小游戏-消灭病毒 [https://github.com/sjatsh/eliminate-virus](https://github.com/sjatsh/eliminate-virus)
