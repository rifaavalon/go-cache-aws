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





MIT License

Copyright (c) 2017 Chris Hickey

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
