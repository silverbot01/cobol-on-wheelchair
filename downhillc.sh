#!/bin/sh

cobc -Wall -x -free cow.cbl cowtemplate.cbl `ls -d controllers/*` -C