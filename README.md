# helloip
A go based docker that runs a web server for load balancer testing

I wrote this based off of the adejonge/helloworld docker image, which is a super tiny container I always use for testing.  But when you run that docker behind a load balancer, or on multiple containers, it always prints the same hello world message.

To rectify this I modified the code to look up the hostname and ip address of eth0, and print them to the web service.

This version of the docker is based off of the Google Cloud Platform container builder service...
