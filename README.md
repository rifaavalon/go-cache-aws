# go-aws-cache 

# Disclaimer 

This project is in the alpha stages. It has been tested with the AWS API so it does work, however it does need some tweaking. 


# What it does 

Using the go-aws-sdk it reaches out to the AWS API and pulls data down into Redis (or ElasticCache using Redis), then spits that data back out as a RESTful API for consumption. 


# What does it support? 

So far only EC2 at the moment, but I would like for it to support more AWS services. 



# How do I use it? 

That's entirely up to you. You can use Task Scheduler if you want to compile it as a Windows binary, cron if you are using Linux for the time being. Later iterations of the code "should just run" at timed intervals, using go's time libraries.

You will want to run this on an AWS instance inside of your environment. As it is designed to use an IAM profile for authentication. 




