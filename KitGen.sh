#!/bin/sh
kit generate service signaling --gorilla
kit generate service signaling --dmw
kit generate service signaling --gorilla
kit generate service signaling -t grpc
