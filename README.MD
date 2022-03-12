# forum-image-upload

## Authors: Zhaksylyk Omizak

## How to run 

Run the following commands:
1. For building an exec file: make build
2. For running the project: make run
3. For creating an image and running container: make docker


## Objectives

You must follow the same [principles](https://github.com/01-edu/public/blob/master/subjects/forum/README.md) as the first subject.

In **forum image upload**, registered users have the possibility to create a post containing an image as well as text.

* When viewing the post, users and guests should see the image associated to it.

There are several extensions for images like: JPEG, SVG, PNG, GIF, etc. In this project you have to handle at least JPEG, PNG and GIF types.

The max size of the images to load should be 20 mb. If there is an attempt to load an image greater than 20mb, an error message should inform the user that the image is too big.

## Hints

* Be cautious with the size of the images.


## Instructions

* The backend must be written in Go.
* You must handle website errors.
* The code must respect the [good practices](https://github.com/01-edu/public/blob/master/subjects/good-practices/README.md)
* It is recommended that the code should present a test file

## Allowed packages

* All [standard go](https://pkg.go.dev/std) packages are allowed.
* github.com/mattn/go-sqlite3
* golang.org/x/crypto/bcrypt
* github.com/satori/go.uuid

### This project will help you learn about:

* Image manipulation
* Image types
