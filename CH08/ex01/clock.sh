#!/bin/bash
TZ=US/Eastern ./myclock -port 8010 &
TZ=Asia/Tokyo ./myclock -port 8020 &
TZ=Europe/London ./myclock -port 8030&
./clockwall/clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

