package main

/*
#cgo LDFLAGS:  -lcurl
#cgo CPPFLAGS: -I/usr/local/opt/curl/include
#include <string.h>
#include <curl/curl.h>

void curl(void) {
    CURL *curl;
    curl = curl_easy_init();
    curl_easy_setopt(curl, CURLOPT_URL, "https://qiita.com/");
    curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, 0);
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
}
*/
import "C"

func main() {
	C.curl()
}
