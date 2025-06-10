package main

import (
	"net/http"
)

func GetHandlerMap() map[string]string {
	handlerMap := make(map[string]string)

	handlerMap[""] = ""
	handlerMap[""] = ""
	handlerMap[""] = ""
	handlerMap[""] = ""

	return handlerMap
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	uriMap := GetHandlerMap()

	return nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
